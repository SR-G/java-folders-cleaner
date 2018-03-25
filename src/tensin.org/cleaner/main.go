package main

import (
	"fmt"
	"os"

	"tensin.org/cleaner/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
