package cmd

import (
	"fmt"

	"github.com/lupinelab/gowake/pkg"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen [macaddress]",
	Short: "Listen for a magic packet",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		remote, macaddr, err := pkg.ListenMagicPacket(port)
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

func init() {
	gowakeCmd.AddCommand(listenCmd)
}
