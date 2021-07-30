package cmd

import (
	"cli/cmd/api"
	"fmt"

	"github.com/spf13/cobra"
)

var projectId string

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
		fmt.Println(api.GetVpcsList(projectId))
	},
}
