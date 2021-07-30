package cmd

import (
	"cli/cmd/api"
	"fmt"

	"github.com/spf13/cobra"
)

var userId string

var projCmd = &cobra.Command{
	Use:   "projs [options]",
	Short: "subcommand for projs",
	Long:  "use another subcommand next",
}

var projList = &cobra.Command{
	Use:   "list",
	Short: "list all projects",
	Long:  "list all projects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(api.GetProjList(userId))
	},
}
