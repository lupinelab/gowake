package cmd

import (
	"fmt"

	"github.com/lupinelab/gowake/pkg"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for a magic packet",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		cont, _ := cmd.Flags().GetBool("continuous")
		fmt.Printf("Listening for magic packets on port %d:\n", port)
		for true {
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
			if !cont {
				break
			}
		}
	},
}

func init() {
	var continuous bool
	listenCmd.Flags().BoolVarP(&continuous, "continuous", "c", false, "Listen continuously for magic packets")
	gowakeCmd.AddCommand(listenCmd)
}
