package gowake

import (
	"fmt"

	gowake "github.com/lupinelab/gowake/internal"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen [macaddress]",
	Short: "Listen for a magic packet",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		remote, macaddr, err := gowake.Listen(port)
		if err != nil {
			if err.Error() == fmt.Sprintf("listen udp 0.0.0.0:%d: bind: permission denied", port) {
				fmt.Println("Please run as elevated user")
				return
			} else {
				fmt.Println(err.Error())
				return
			}
		}
		fmt.Printf("%v from %v\n", macaddr, remote.String())
	},
}
