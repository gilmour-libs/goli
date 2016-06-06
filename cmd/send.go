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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	ui "github.com/meson10/goerrui"
	G "gopkg.in/gilmour-libs/gilmour-e-go.v4"
)

var file string

func contentAndTopic(args []string) (content string, topic string) {
	if len(args) < 1 || len(args) > 2 {
		return
	}

	topic = args[0]

	if file != "" {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			ui.Alert("Input file %v not found on Disk", file)
			return
		}

		if data, err := ioutil.ReadFile(file); err != nil {
			ui.Alert("Error %v while parsing File", err.Error(), file)
			return
		} else {
			content = string(data)
		}
	} else if len(args) == 2 {
		content = args[1]
	} else {
		ui.Alert("Must provide content as position argunent of a file")

		return
	}

	return
}

// sendCmd respresents the send command
var sendCmd = &cobra.Command{
	Use:   "send <topic> <json_string> OR -f path/to/json/file",
	Short: "Send data to request/slot topic",
	Run: func(cmd *cobra.Command, args []string) {
		content, topic := contentAndTopic(args)
		if content == "" || topic == "" {
			cmd.Help()
			return
		}

		engine := getEngine()
		defer engine.Stop()
		engine.Start()

		data := G.NewMessage().SetData(content)

		if isSlot {
			if _, err := engine.Signal(topic, data); err != nil {
				ui.Alert(err.Error())
			}
		} else {
			req := engine.NewRequest(topic)
			resp, err := req.Execute(data)
			if err != nil {
				ui.Alert(err.Error())
				return
			}

			var output string
			if err := resp.Next().GetData(&output); err != nil {
				ui.Alert(err.Error())
				return
			} else {
				fmt.Println(output)
				return
			}
		}

		return
	},
}

func init() {
	sendCmd.Flags().BoolVarP(&isSlot, "slot", "", false, "Is this topic a slot?")
	sendCmd.Flags().StringVarP(&file, "file", "f", "", "File with Content")
	RootCmd.AddCommand(sendCmd)
}
