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
	"sync"

	"github.com/spf13/cobra"
	G "gopkg.in/gilmour-libs/gilmour-e-go.v4"
)

var isSlot bool
var isRequest bool

// tailCmd respresents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Listen to messages that arrive on a Gilmour Topic",
	Long: `Keep listening to messages that arrive on a topic.

	All gilmour slots and Requests emit Signals on:
	gilmour.slot.log.gilmour.<slot|request>.<topic>

	A signal is sent as early as a message is received for that topic, and in case
	of request the response will also be emitted.

	You must specify the --slot or --request switch as well to define wether goli
	should sniff a slot or a request.

	Usage:
	goli tail <topic>`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		engine := getEngine()

		topic := args[0]
		switch {
		case isSlot:
			topic = fmt.Sprintf("gilmour.slot.%v", topic)
		case isRequest:
			topic = fmt.Sprintf("gilmour.request.%v", topic)
		default:
			cmd.Help()
			return
		}

		topic = fmt.Sprintf("gilmour.log.%v", topic)
		fmt.Println("Starting tail on", topic)

		engine.Slot(topic, func(req *G.Request) {
			var msg interface{}
			if err := req.Data(&msg); err != nil {
				log.Println("Cannot parse log %v", err.Error())
				return
			}

			log.Println(req.Sender(), "->", msg)
		}, nil)

		engine.Start()

		var wg sync.WaitGroup
		wg.Add(1)
		wg.Wait()
	},
}

func init() {
	tailCmd.Flags().BoolVarP(&isSlot, "slot", "", false, "Is this topic a slot?")
	tailCmd.Flags().BoolVarP(&isRequest, "request", "", false, "Is this topic a Request?")

	RootCmd.AddCommand(tailCmd)
}
