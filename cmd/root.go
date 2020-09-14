package cmd

import (
	"os"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/winkoz/plonk/internal/io/logger"
)

var rootCmd = &cobra.Command{
	Use:   "plonk",
	Short: "Plonk is a deploy manager for kubernetes apps using kubeclt",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("RUN COMMAND - Verbosity: %s", logger.Severity)
		// Do Stuff Here
	},
}

// Execute executes command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)

		os.Exit(1)
	}
}
