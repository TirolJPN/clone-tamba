package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func fetchCmd() *cobra.Command {
cmd := &cobra.Command{
		Use: "fetch",
		Short: "fetch is command to fetch MySQL data",
		// RangeArgs(min, max) - the command will report an error if the number of args is not between the minimum and maximum number of expected args.
		Args: cobra.RangeArgs(2,2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			itemOne, err := strconv.Atoi(args[0])

			if err != nil {
				log.Fatal(err)
				return nil
			}

			itemTwo, err := strconv.Atoi(args[1])

			if err != nil {
				log.Fatal(err)
				return nil
			}

			fmt.Println(itemOne + itemTwo)

			return nil
		},
	}

	return cmd
}