package cmd

import (
	"github.com/spf13/cobra"
	"github.com/winkoz/plonk/internal/io/log"
)

// AddVerbosity adds a --verbose and --quiet flag to the command and all subcommands.
func AddVerbosity(cmd *cobra.Command) {
	verboseFlag := cmd.PersistentFlags().VarPF(&log.Severity, "verbose", "v", "More verbose output")
	quiteFlag := cmd.PersistentFlags().VarPF(&log.Severity, "quiet", "q", "Less verbose output")

	// Allow using the flags without arguments, since we don't want an argument in the first place. This makes `sc -q -q -q` work.
	verboseFlag.NoOptDefVal = "DEBUG"
	quiteFlag.NoOptDefVal = "INFO"
}
