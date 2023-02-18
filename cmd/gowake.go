package cmd

import (
	"fmt"
	"regexp"
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

		// Get IP
		ip, _ := cmd.Flags().GetString("ip")

		is_ip, _ := regexp.MatchString(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`, ip)
		if !is_ip {
			fmt.Println("Please provide a valid IP address")
			return
		}

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

		sendoptions := pkg.SendOptions{Mp: mp, Port: port, IP: ip}

		// Send packet
		pkg.SendMagicPacket(sendoptions)
		fmt.Printf("Sent magic packet to %v\n", args[0])
	},
}

func Execute() error {
	return gowakeCmd.Execute()
}

func init() {
	var port int
	var ip string
	gowakeCmd.PersistentFlags().IntVarP(&port, "port", "p", 9, "Port to send or listen for magic packet")
	gowakeCmd.Flags().StringVarP(&ip, "ip", "i", "255.255.255.255", "Destination (IP or broadcast address) for the magic packet")
	gowakeCmd.PersistentFlags().BoolP("help", "h", false, "Print help")
	gowakeCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
	cobra.EnableCommandSorting = false
	gowakeCmd.CompletionOptions.DisableDefaultCmd = true
}
