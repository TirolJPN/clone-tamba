package main 

import (
	"os"
	"fmt"
    "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Name = "sampleApp"
  app.Usage = "This app echo input arguments"
  app.Version = "0.0.1"

  app.Action = func (context *cli.Context) error {
    fmt.Println(context.Args().Get(0))
    return nil
  }

  app.Run(os.Args)
  
}