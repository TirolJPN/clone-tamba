package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "root command for all commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		fetchCmd(),
	)
}