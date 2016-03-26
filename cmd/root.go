package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "goli",
	Short: "Goli is a fast debugger for Gilmour applications",
	Long:  `A Fast and Powerful debugger & monitor for your Gilmour architecture.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
