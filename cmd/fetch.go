package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/TirolJPN/clone-tamba/sql/file"
)

func fetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "fetch",
		Short: "fetch is command to fetch MySQL data",
		// RangeArgs(min, max) - the command will report an error if the number of args is not between the minimum and maximum number of expected args.
		Args: cobra.RangeArgs(1,100),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				print("An argument is needed at least to fetch problem information by problem id")
				return nil
			}
			for _, elem := range args {
				print(elem)
				file.StatusAndFilePath(elem)
			}
			return nil
		},
	}

	return cmd
}