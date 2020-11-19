package cmd

import "github.com/spf13/cobra"

const (
	defaultDeployEnvironment = "production"
)

// CobraHandler handler interface for a cobra command
type CobraHandler func(cmd *cobra.Command, args []string)
