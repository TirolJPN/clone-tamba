package main 

import (
	"os"
	"fmt"
    "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Name = "clone-tamba"
  app.Usage = "This app is clone of tamba system"
  app.Version = "0.0.1"

  app.Action = func (context *cli.Context) error {
    fmt.Println("An argument is needed.")
    return nil
  }

  app.Run(os.Args)
  
}