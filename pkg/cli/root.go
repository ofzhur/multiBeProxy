package cli

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "multiBeProxy",
		Short: "Start server",
		Long:  "Multi-target http proxy with mirroring requests to all backends",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
