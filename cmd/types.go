package cmd

import "github.com/spf13/cobra"

type CobraHandler func(cmd *cobra.Command, args []string)

// Constants
const defaultTemplatesPath = "./templates"
const defaultCustomTemplatesPath = "$HOME/plonk/templates"
const deployFolderName = "deploy"
const deployVariablesPath = deployFolderName + "/variables"
