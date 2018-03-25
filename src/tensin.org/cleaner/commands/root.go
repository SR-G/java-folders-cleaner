package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ConfigurationFileName is the configuration file name that will be used
var ConfigurationFileName string

// RootCmd is the main command = the program itself
var RootCmd = &cobra.Command{
	Use:   "java-folders-cleaner",
	Short: "java-folders-cleaner",
	Long:  `java-folders-cleaner`,
}

func init() {
	// RootCmd.PersistentFlags().StringVarP(&ConfigurationFileName, "configuration", "", "watchthatpage.json", "Configuration file name. Default is binary name + .json (e.g. 'watchthatpage.json'), in the same folder than the binary itself")
	// daemon.Server.Start()
	fmt.Println("start")
}