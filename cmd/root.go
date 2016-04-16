package cmd

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/cobra"
	G "gopkg.in/gilmour-libs/gilmour-e-go.v4"
	"gopkg.in/gilmour-libs/gilmour-e-go.v4/backends"
)

var redisPort int
var once sync.Once
var engine *G.Gilmour
var redisHost, redisPass string

var RootCmd = &cobra.Command{
	Use:   "goli",
	Short: "Goli is a fast debugger for Gilmour applications",
	Long:  `A Fast and Powerful debugger & monitor for your Gilmour architecture.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	RootCmd.PersistentFlags().StringVarP(
		&redisHost, "host", "", "127.0.0.1", "Redis host to conenct to",
	)

	RootCmd.PersistentFlags().IntVarP(
		&redisPort, "port", "", 6379, "Redis Port to connect to",
	)

	RootCmd.PersistentFlags().StringVarP(
		&redisPass, "password", "", "", "Password to connect to Redis",
	)
}

func getEngine() *G.Gilmour {
	once.Do(func() {
		redis := backends.MakeRedis(
			fmt.Sprintf("%v:%v", redisHost, redisPort),
			redisPass,
		)

		engine = G.Get(redis)
	})

	return engine
}
