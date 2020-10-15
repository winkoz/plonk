package cmd

import (
	"github.com/spf13/cobra"
	"github.com/winkoz/plonk/internal/io/log"
)

// AddVerbosity adds a --verbosity flag
func addVerbosity(cmd *cobra.Command) {
	verbosityFlag := cmd.PersistentFlags().VarPF(&log.Severity, "verbosity", "v", "Define verbosity level [debug, info, warn, error, fatal]")

	verbosityFlag.DefValue = "info"
	verbosityFlag.Value.Set("info")
}
