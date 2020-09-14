package cmd

import (
	"fmt"
	"os"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plonk",
	Short: "Plonk is a deploy manager for kubernetes apps using kubeclt",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RUN COMMAND")
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
