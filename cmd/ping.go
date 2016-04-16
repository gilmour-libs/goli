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

// pingCmd respresents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Is the gilmour subscriber still active?",
	Long: `Is the gilmour server stil alive for the given health-ident?
	Heartbeat is subject to a timeout. Read documentation for details.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("ping called")
	},
}

func init() {
	RootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings

	// Cobra supports Persistent Flags which will work for this command and all subcommands
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle" )

}