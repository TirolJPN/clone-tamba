package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/TirolJPN/clone-tamba/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
