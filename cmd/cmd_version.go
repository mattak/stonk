package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
var (
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long:  `Show version`,
		Example: ` stonk version`,
		Run: runCommandVersion,
	}
)

func runCommandVersion(cmd *cobra.Command, args []string) {
	fmt.Println(VERSION)
}
