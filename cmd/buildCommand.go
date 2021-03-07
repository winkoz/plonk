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
	"github.com/winkoz/plonk/internal/building"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io/log"
)

func addBuildCommand(rootCmd *cobra.Command) {
	var buildCmd = &cobra.Command{
		Use:     "build",
		Short:   "Builds a docker project.",
		Run:     newBuildCommandHandler(),
		Example: "plonk build",
		Args:    cobra.MaximumNArgs(1),
	}

	rootCmd.AddCommand(buildCmd)
}

func newBuildCommandHandler() CobraHandler {
	return func(cmd *cobra.Command, args []string) {
		ctx, err := config.NewContextFromFile()
		if err != nil {
			log.Fatal(err)
		}

		env := defaultDeployEnvironment
		if len(args) == 1 {
			env = args[0]
		}

		builder := building.NewBuilder(ctx)

		tag, err := builder.Build(env)
		if err != nil {
			log.Errorf("Failed building current docker project %s.", env, err)
			return
		}
		log.Infof("Build tag: %s", tag)
		log.Info("Build executed successfully.")
	}
}
