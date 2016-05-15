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
	"time"

	"github.com/spf13/cobra"
	G "gopkg.in/gilmour-libs/gilmour-e-go.v4"
)

var wait = 1000
var count int
var waitTime = 5

// pingCmd respresents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Is the gilmour subscriber still active?",
	Long: `Is the gilmour server stil alive for the given health-ident?
	Heartbeat is subject to a timeout. Read documentation for details.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}

		host := args[0]
		engine := getEngine()
		defer engine.Stop()
		engine.Start()

		data := G.NewMessage().SetData("ping?")

		for i := 0; i <= count; {
			if count > 0 {
				i += 1
			}

			opts := G.NewRequestOpts().SetTimeout(waitTime)
			req := engine.NewRequestWithOpts(
				fmt.Sprintf("gilmour.health.%v", host),
				opts,
			)

			now := time.Now().UTC()
			resp, err := req.Execute(data)
			if err != nil {
				fmt.Println(fmt.Sprintf("Destination host %v unreachable", host))
				continue
			}

			elapsed := time.Since(now)
			msg := resp.Next()
			if resp.Code() != 200 {
				var x string
				msg.GetData(&x)
				fmt.Println(fmt.Sprintf("Error %v contacting host %v", x, host))
			} else {
				fmt.Println(fmt.Sprintf("Success from %v time=%v", host, elapsed))
			}

			time.Sleep(time.Duration(wait) * time.Second)
		}
	},
}

func init() {
	pingCmd.Flags().IntVarP(&count, "count", "c", count, "Stop after sending (and receiving) count ECHO_RESPONSE packets. If this option is not specified, ping will operate until interrupted.")

	pingCmd.Flags().IntVarP(&waitTime, "waittime", "w", waitTime, "Time in seconds to wait for a reply for each packet sent.  If a reply arrives later, the packet is not printed as replied, but considered as replied when calculating statistics.")

	pingCmd.Flags().IntVarP(&wait, "wait", "i", wait, "Wait wait seconds between sending each packet. The default is to wait for 1000ms between each packet.  The wait time may be as low as 1ms, but only sensible-users would specify values less than 10ms.")

	RootCmd.AddCommand(pingCmd)
}
