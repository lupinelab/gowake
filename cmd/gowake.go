package cmd

import (
	"fmt"
	"github.com/lupinelab/gowake/pkg"
	"github.com/spf13/cobra"
)

var gowakeCmd = &cobra.Command{
	Use:   "gowake [macaddress]",
	Short: "Send a magic packet in order wake a compliant machine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get port
		port, _ := cmd.Flags().GetInt("port")

		// Check for mac address
		if len(args) < 1 {
			fmt.Println("Please provide a MAC address")
			return
		}

		// Build packet
		mp, err := pkg.BuildMagicPacket(args[0])
		if err != nil {
			return
		}

		// Send packet
		pkg.SendMagicPacket(mp, port)
		fmt.Printf("Sent magic packet to %v\n", args[0])
	},
}

func Execute() error {
	return gowakeCmd.Execute()
}

func init() {
	var port int
	gowakeCmd.PersistentFlags().IntVarP(&port, "port", "p", 9, "Port on which send or listen for magic packet")
	gowakeCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	gowakeCmd.PersistentFlags().Lookup("help").Hidden = true
	cobra.EnableCommandSorting = false
	gowakeCmd.CompletionOptions.DisableDefaultCmd = true
}
