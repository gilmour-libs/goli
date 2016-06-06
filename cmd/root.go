package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/spf13/cobra"
	G "gopkg.in/gilmour-libs/gilmour-e-go.v4"
	"gopkg.in/gilmour-libs/gilmour-e-go.v4/backends"
	"gopkg.in/meson10/goraph.v0"
)

var redisPort int
var engineOnce, backendOnce sync.Once
var backend *backends.Redis
var engine *G.Gilmour
var redisHost, redisPass string

var RootCmd = &cobra.Command{
	Use:   "goli",
	Short: "Goli is a fast debugger for Gilmour applications",
	Long:  `A Fast and Powerful debugger & monitor for your Gilmour architecture.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmd.Help()
	},
}

func trapInterrupt(engine *G.Gilmour) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(e *G.Gilmour) {
		for sig := range c {
			// sig is a ^C, handle it
			log.Println("Caught interrupt", sig)
			e.Stop()
			os.Exit(0)
		}
	}(engine)
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

	RootCmd.PersistentFlags().BoolVarP(
		&goraph.MachineReadable, "machine-readable", "", false, "Use comma separated graph output to be used for Command line tools",
	)
}

func getBackend() *backends.Redis {
	backendOnce.Do(func() {
		backend = backends.MakeRedis(
			fmt.Sprintf("%v:%v", redisHost, redisPort),
			redisPass,
		)
	})
	return backend
}

func getEngine() *G.Gilmour {
	engineOnce.Do(func() {
		redis := getBackend()
		engine = G.Get(redis)
	})

	return engine
}
