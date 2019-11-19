package file

import (
	"database/sql"
	"fmt"
	"os"
)

func All() {
	// usage: cnn := sql.Open("mysql", "user:password@tcp(host:port)/dbname")
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := "SELECT * FROM Problem"
	rows, err := db.Query(query)
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
}

// 引数で指定されたproblem_idのすべての解答について，submission_id, timestamp, ローカルにあるファイルのパスを[]interface{}で返す
func StatusAndFilePath(problem_id string) [][]string {
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT file_name, submission_id, timestamp FROM File WHERE problem_id = \"%s\"", problem_id)
	rows, err := db.Query(query)

	columns, err := rows.Columns() // fetch culmn names
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	fetchList := [][]string{}
	for rows.Next() {
		tmpFetch := []string{}
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for _, col := range values {
			// Here we can check if the value is NULL value
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			tmpFetch = append(tmpFetch, value)
		}
		fetchList = append(fetchList, tmpFetch)
	}
	return fetchList
}