// Copyright ©2016 NAME HERE <EMAIL ADDRESS>
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

// chckCmd respresents the chck command
var chckCmd = &cobra.Command{
	Use:   "chck",
	Short: "Check if there are active listeners for a topic",
	Long: `From an architect's view, are there atleast one active listener(s)
	that can serve this topic, subject to a timeout and return error otherwise.
	For example:

	For a rqeuest topic item.purchase.520, are there any liteners in the current
	setup that can serve this request`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("chck called")
	},
}

func init() {
	RootCmd.AddCommand(chckCmd)

	// Here you will define your flags and configuration settings

	// Cobra supports Persistent Flags which will work for this command and all subcommands
	// chckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly
	// chckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle" )

}
