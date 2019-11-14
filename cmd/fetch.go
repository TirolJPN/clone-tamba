package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func fetchCmd() *cobra.Command {
cmd := &cobra.Command{
		Use: "fetch",
		Short: "fetch is command to fetch MySQL data",
		// RangeArgs(min, max) - the command will report an error if the number of args is not between the minimum and maximum number of expected args.
		Args: cobra.RangeArgs(0,0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// usage: cnn := sql.Open("mysql", "user:password@tcp(host:port)/dbname")
			dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
			db, err := sql.Open("mysql", dataSourceName)
			if err != nil {
				panic(err.Error())
			}
			defer db.Close()

			rows, err := db.Query("SELECT * FROM Problem")
			if err != nil {
				panic(err.Error())
			}

			columns, err := rows.Columns() // fetch culmn names
			if err != nil {
				panic(err.Error())
			}

			values := make([]sql.RawBytes, len(columns))

			scanArgs := make([]interface{}, len(values))
			for i := range values {
				scanArgs[i] = &values[i]
			}

			for rows.Next() {
				err = rows.Scan(scanArgs...)
				if err != nil {
					panic(err.Error())
				}

				var value string
				for i, col := range values {
					// Here we can check if the value is NULL value
					if col == nil {
						value = "NULL"
					} else {
						value = string(col)
					}
					fmt.Println(columns[i], ":", value)
				}
				fmt.Println("-------------------------------------------")
			}
			return nil
		},
	}

	return cmd
}