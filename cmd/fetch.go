package cmd

import (
	"fmt"
	"github.com/TirolJPN/clone-tamba/gragh"
	"github.com/TirolJPN/clone-tamba/sql/file"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)


func fetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "fetch",
		Short: "fetch is command to fetch MySQL data",
		// RangeArgs(min, max) - the command will report an error if the number of args is not between the minimum and maximum number of expected args.
		Args: cobra.RangeArgs(1,100),
		RunE: func(cmd *cobra.Command, args []string) error {
			var nodes gragh.Nodes

			if len(args) == 0 {
				print("An argument is needed at least to fetch problem information by problem id")
				return nil
			}
			for _, problemId := range args {
				baseDir := os.Getenv("PATH_FROM_WORKSPACE_TO_CLUSTERING_RESULT") +
							problemId +
							"/cosine/complete/result/"
				filePaths := dirWalk(baseDir)

				fetchList := file.StatusAndFilePath(problemId)
				for _, culumn := range fetchList {
					fileName := culumn[0]
					submissionId := culumn[1]
					timestamp := culumn[2]
					filePath, lexicalIndex, metricalIndex, err :=  searchFilePath(filePaths, fileName)
					if err != nil {
						fmt.Printf("Not found :%s %s\n", fileName, submissionId)
					}else {
						tmp := gragh.NewNode(submissionId, fileName, filePath, timestamp, lexicalIndex, metricalIndex)
						println(tmp)
						nodes = append(nodes, gragh.NewNode(submissionId, fileName, filePath, timestamp, lexicalIndex, metricalIndex))
					}

				}
				// process to make directed graph by timestamps
			}

			// for tests
			for _, elem := range nodes {
				fmt.Println(elem)
			}
			return nil
		},
	}
	return cmd
}

// ファイル名を引数にして，env情報をもとにファイル検索を行い，ファイルの絶対パスを表す文字列を返す
func searchFilePath(filePaths []string, fileName string) (filePath string, lexicalIndex int, metriaclIndex int, err error ){
	for _, filePath := range filePaths {
		sliced := strings.Split(filePath, "\\")
		lexicalIndex, _ := strconv.Atoi(sliced[len(sliced) - 3])
		metriaclIndex, _ := strconv.Atoi(sliced[len(sliced) - 2])
		targetFileName := sliced[len(sliced) - 1]
		if targetFileName == fileName {
			return filePath, lexicalIndex, metriaclIndex, nil
		}
	}
	return "", -1, -1, os.ErrExist
}

// 引数に与えたパス以下のすべてのフォルダ・ファイルを列挙し，リストに格納する
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











