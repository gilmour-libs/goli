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
	"io/ioutil"
	"log"
	"os"

	ui "github.com/meson10/goerrui"
	"github.com/spf13/cobra"
)

var file string

// rqstCmd respresents the rqst command
var rqstCmd = &cobra.Command{
	Use:   "rqst",
	Short: "Send JSON request and publish the response",
	Long: `Send JSON data to a gilmour topic and wait for the response which can
	later be piped to an external program`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		var json string

		if file != "" {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				ui.Alert("JSON input file %v not found on Disk", file)
				return
			}

			if data, err := ioutil.ReadFile(file); err != nil {
				ui.Alert("Error %v while parsing JSON input", err.Error(), file)
				return
			} else {
				json = string(data)
			}
		} else if len(args) == 1 {
			json = args[0]
		} else {
			ui.Alert("Provide one of the JSON inputs")
			cmd.Help()
			return
		}

		log.Println(json)
	},
}

func init() {
	rqstCmd.Flags().StringVarP(&file, "file", "f", "", "File with JSON data")
	RootCmd.AddCommand(rqstCmd)
}
