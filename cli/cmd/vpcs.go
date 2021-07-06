package cmd

import (
	"fmt"
	"strings"
	
	"github.com/spf13/cobra"
)

var vpcs = &cobra.Command{
	Use:	"vpcs [options]",
	Short: "subcommand for vpcs",
	Long: "use another subcommand next",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))
	},
}