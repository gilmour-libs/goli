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

	"github.com/spf13/cobra"
)

// rqstCmd respresents the rqst command
var rqstCmd = &cobra.Command{
	Use:   "rqst",
	Short: "Send JSON request and publish the response",
	Long: `Send JSON data to a gilmour topic and wait for the response which can
	later be piped to an external program`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("rqst called")
	},
}

func init() {
	RootCmd.AddCommand(rqstCmd)

	// Here you will define your flags and configuration settings

	// Cobra supports Persistent Flags which will work for this command and all subcommands
	// rqstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly
	// rqstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle" )

}
