package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var vpcsCmd = &cobra.Command{
	Use:   "vpcs [options]",
	Short: "subcommand for vpcs",
	Long:  "use another subcommand next",
}

var vpcsList = &cobra.Command{
	Use:   "list",
	Short: "list all vpcs",
	Long:  "list all vpcs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))
	},
}
