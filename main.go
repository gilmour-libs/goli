package main

import (
	"fmt"
	"os"

	"gopkg.in/gilmour-libs/goli.v1/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
