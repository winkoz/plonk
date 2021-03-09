package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/winkoz/plonk/internal/io/log"
)

var rootCmd = &cobra.Command{
	Use:   "plonk",
	Short: "Plonk is a deploy manager for kubernetes apps using kubeclt",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("RUN COMMAND - Verbosity: %s", log.Severity)
		// Do Stuff Here
	},
}

// Execute executes command
func Execute() {
	rootCmd := newRootCommand()

	// Logger Verbosity Configuration
	addVerbosity(rootCmd)

	// Commands
	addInitCommand(rootCmd)
	addDeployCommand(rootCmd)
	addDestroyCommand(rootCmd)
	addDiffCommand(rootCmd)
	addShowCommand(rootCmd)
	addLogsCommand(rootCmd)
	addRestartCommand(rootCmd)

	// Execute
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "plonk",
		Short: "Plonk is a deploy manager for kubernetes apps using kubeclt",
		Long:  "",
		Run:   newRootCommandHandler(),
	}
}

func newRootCommandHandler() CobraHandler {
	return func(cmd *cobra.Command, args []string) {
		log.Debugf("RUN COMMAND - Verbosity: %s", log.Severity)
	}
}
