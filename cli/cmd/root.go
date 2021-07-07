package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultProjectId = "0b5a73ddd98027372f2ec00668b88856"
	defaultUserId    = "0ce43a5b788024e71f03c0060aaf6125"
)

var Number int
var rootCmd = &cobra.Command{
	Use:   "sbcli [subcommands]",
	Short: "Sbercloud cli",
	Long: `main command for
	Sbercloud cli`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(vpcsCmd, projCmd)
	vpcsCmd.AddCommand(vpcsList)
	projCmd.AddCommand(projList)
	vpcsList.Flags().IntVarP(&Number, "number", "n", 2, "limit of vpcs shown")
	vpcsList.Flags().StringVarP(&projectId, "pid", "p", defaultProjectId, "project id for query")
	projList.Flags().StringVarP(&userId, "uid", "u", defaultUserId, "user id for the query")
}
