// Copyright ©2016 Piyush Verma <piyush@piyushverma.net>
//
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package cmd

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	G "gopkg.in/gilmour-libs/gilmour-e-go.v4"
	"gopkg.in/meson10/goraph.v0"
)

func testHost(engine *G.Gilmour, host string) ([]string, error) {
	data := G.NewMessage().SetData("ping?")
	req := engine.NewRequest(fmt.Sprintf("gilmour.health.%v", host))

	var topics []string
	resp, err := req.Execute(data)
	if err != nil {
		return topics, err
	}

	msg := resp.Next()
	msg.GetData(&topics)
	return topics, nil
}

func makeGraph(engine *G.Gilmour, idents map[string]string) goraph.Graph {
	root := goraph.MakeGraph()

	var wg sync.WaitGroup

	for host, _ := range idents {
		wg.Add(1)

		go func(root goraph.Graph, engine *G.Gilmour, host string) {
			defer wg.Done()
			node := goraph.NewNode(host)
			node.SetType("host")

			topics, err := testHost(engine, host)
			if err != nil {
				msg := fmt.Sprintf("Error %v communicating with %v", host, err.Error())
				node.SetDirty()
				node.AddNote(msg)
				return
			}

			for _, t := range topics {
				tnode := goraph.NewNode(t)
				tnode.SetType("topic")
				node.AddChild(tnode)
			}

			root.AddChild(node)
		}(root, engine, host)
	}

	wg.Wait()
	return root
}

func globIdent(idents map[string]string, patterns []string) map[string]string {
	if len(patterns) == 0 {
		return idents
	}

	res := map[string]string{}

	for _, i := range patterns {
		if strings.HasSuffix(i, "*") {
			ident := strings.Split(i, "*")[0]
			for k, _ := range idents {
				if strings.HasPrefix(k, ident) {
					res[k] = "true"
				}
			}
		} else if _, ok := idents[i]; ok {
			res[i] = "true"
		}
	}

	return res
}

// statCmd respresents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Get all registered health-idents",
	Long: `Generates a Tree of Hosts and the topics they are listening to:

	By default, it pings all hosts registered with Gilmour for health checks.
	You can specify host(s) as arguments to stat command.

	You can also provide pattern(s), but the pattern must end with an asterisk.

	Example Usage:

	piyush:goli master λ goli stat SatellitePro.local* hello*
	root
	└─── SatellitePro.local-pid-88830-uuid-38ec66c0-a09a-4005-be7e-1d2a4b0adb00
	│    ├─── gilmour.request.request.manager.bringup.cassandra
	│    ├─── gilmour.request.request.manager.reconfigure.cassandra
	│    ├─── request.manager.addnodes.cassandra
	│    ├─── gilmour.request.request.manager.addnodes.kafka
	│    ├─── request.manager.spark.info
	│    ├─── gilmour.request.request.manager.zookeeper.removenodes
	│    ├─── request.manager.teardown.cassandra
	│    ├─── request.manager.addusers.cassandra
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		redis := getBackend()
		engine := getEngine()
		defer engine.Stop()

		active_idents, err := redis.ActiveIdents()
		if err != nil {
			log.Println("Cannot fetch Idents:", err.Error())
			return
		}

		idents := globIdent(active_idents, args)

		engine.Start()
		graph := makeGraph(engine, idents)
		graph.Walk()
	},
}

func init() {
	RootCmd.AddCommand(statCmd)
}
