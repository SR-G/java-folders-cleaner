package commands

import (
	"fmt"

	"tensin.org/cleaner/core/version"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the chainer command.",
	Long:  `Prints the version of the chainer command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(core.Version)
	},
}
