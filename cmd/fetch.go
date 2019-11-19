package cmd

import (
	"github.com/TirolJPN/clone-tamba/sql/file"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
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
			for _, problemId := range args {
				baseDir := os.Getenv("PATH_FROM_WORKSPACE_TO_CLUSTERING_RESULT") +
							problemId +
							"/cosine/complete/result/"
				filePaths := dirWalk(baseDir)
				for _, tmp := range(filePaths) {
					println(tmp)
				}

				fetchList := file.StatusAndFilePath(problemId)
				for _, culumn := range fetchList {
					fileName := culumn[0]
					submissionId := culumn[1]
					timeStamp := culumn[2]

					println(fileName, submissionId, timeStamp)
				}


				// process to make directed graph by timestamp
			}
			return nil
		},
	}
	return cmd
}

// ファイル名を引数にして，env情報をもとにファイル検索を行い，ファイルの絶対パスを表す文字列を返す
func searchFilePath() {}

func dirWalk(dir string ) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirWalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}











