package cmd

import "github.com/spf13/cobra"

type CobraHandler func(cmd *cobra.Command, args []string)
