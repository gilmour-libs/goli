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

// rqstCmd respresents the rqst command
var rqstCmd = &cobra.Command{
	Use:   "rqst <topic>",
	Short: "Send request data and publish the response",
	Long: `Send request data to a gilmour topic and wait for the response which can
	later be piped to an external program`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		var content string

		if len(args) < 1 || len(args) > 2 {
			cmd.Help()
			return
		}

		engine := getEngine()
		defer engine.Stop()
		engine.Start()

		topic := args[0]

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
			cmd.Help()
			return
		}

		req := engine.NewRequest(topic)
		resp, err := req.Execute(G.NewMessage().SetData(content))
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
	},
}

func init() {
	rqstCmd.Flags().StringVarP(&file, "file", "f", "", "File with Content")
	RootCmd.AddCommand(rqstCmd)
}
