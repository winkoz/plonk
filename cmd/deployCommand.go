package cmd

/*
Copyright © 2020 Winkoz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"github.com/spf13/cobra"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/deployment"
	"github.com/winkoz/plonk/internal/io/log"
)

func addDeployCommand(rootCmd *cobra.Command, ctx config.Context) {
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a project deployable by plonk.",
		Long: `This triggers a questionnaire for the basic information of a project
		and then generates the deploy folder with the basic project files.
		`,
		Run:     newInitCommandHandler(ctx),
		Example: "plonk init",
	}
	rootCmd.AddCommand(initCmd)
}

func newDeployCommandHandler(ctx config.Context) CobraHandler {
	return func(cmd *cobra.Command, args []string) {

		env, err := cmd.PersistentFlags().GetString(cmdFlagEnvironment)
		if err != nil {
			log.Fatalf("Can't execute deploy without: %s", cmdFlagEnvironment)
		}

		d := deployment.NewDeployer(ctx)
		d.Execute(ctx, env)
	}
}
