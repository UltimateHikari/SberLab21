package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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

func init(){
	rootCmd.AddCommand(vpcs)
}
