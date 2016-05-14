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

	"github.com/meson10/goraph"
	"github.com/spf13/cobra"
	G "gopkg.in/gilmour-libs/gilmour-e-go.v4"
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

	for host, _ := range idents {
		node := goraph.NewNode(host)
		node.SetType("host")

		topics, err := testHost(engine, host)
		if err != nil {
			log.Println("Error communicating with", host, err.Error())
			continue
		}

		for _, t := range topics {
			tnode := goraph.NewNode(t)
			tnode.SetType("topic")
			node.AddChild(tnode)
		}

		root.AddChild(node)
	}
	return root
}

// statCmd respresents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Get all registered health-idents",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		redis := getBackend()
		engine := getEngine()
		defer engine.Stop()

		idents, err := redis.ActiveIdents()
		if err != nil {
			log.Println("Cannot fetch Idents:", err.Error())
			return
		}

		engine.Start()
		graph := makeGraph(engine, idents)
		graph.Walk()
	},
}

func init() {
	RootCmd.AddCommand(statCmd)

	// Here you will define your flags and configuration settings

	// Cobra supports Persistent Flags which will work for this command and all subcommands
	// statCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly
	// statCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle" )

}
