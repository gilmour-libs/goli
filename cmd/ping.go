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

	"github.com/meson10/goraph"
	"github.com/spf13/cobra"

	G "gopkg.in/gilmour-libs/gilmour-e-go.v4"
)

func healthGraph(engine *G.Gilmour, idents map[string]string) goraph.Graph {
	root := goraph.MakeGraph()

	for host, _ := range idents {
		node := goraph.NewNode(host)
		node.SetType("host")

		if _, err := testHost(engine, host); err != nil {
			msg := fmt.Sprintf("Error %v communicating with %v", host, err.Error())
			node.SetDirty()
			node.AddNote(msg)
		}

		root.AddChild(node)
	}
	return root
}

// pingCmd respresents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Is the gilmour subscriber still active?",
	Long: `Is the gilmour server stil alive for the given health-ident?
	Heartbeat is subject to a timeout. Read documentation for details.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		redis := getBackend()
		engine := getEngine()
		defer engine.Stop()

		active_idents, err := redis.ActiveIdents()
		if err != nil {
			log.Println("Cannot fetch Idents:", err.Error())
			return
		}

		idents := map[string]string{}

		for _, i := range args {
			if strings.HasSuffix(i, "*") {
				ident := strings.Split(i, "*")[0]
				for k, _ := range active_idents {
					if strings.HasPrefix(k, ident) {
						idents[k] = "true"
					}
				}
			} else if _, ok := active_idents[i]; ok {
				idents[i] = "true"
			}
		}

		engine.Start()
		healthGraph(engine, idents).Walk()
	},
}

func init() {
	RootCmd.AddCommand(pingCmd)
}
