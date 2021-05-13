package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "stonk",
		Short: "Stock market related tools",
		Long:  `Stock market related tools`,
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(SymbolCmd)
	RootCmd.AddCommand(PriceCmd)
}
