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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/winkoz/plonk/internal/building"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/deployment"
	"github.com/winkoz/plonk/internal/io/log"
)

func addDeployCommand(rootCmd *cobra.Command) {
	var deployCmd = &cobra.Command{
		Use:     "deploy",
		Short:   "Deploys a project previously initialised by plonk.",
		Run:     newDeployCommandHandler(),
		Example: "plonk deploy",
		Args:    cobra.MaximumNArgs(1),
	}

	skipBuild := false
	deployCmd.Flags().BoolVarP(&skipBuild, "skip-build-n-publish", "", false, "skip docker build step")

	skipCache := true
	deployCmd.Flags().BoolVarP(&skipCache, "skip-cache", "", true, "skip docker cache when building image")

	rootCmd.AddCommand(deployCmd)
}

func newDeployCommandHandler() CobraHandler {
	return func(cmd *cobra.Command, args []string) {
		ctx, err := config.NewContextFromFile()
		if err != nil {
			log.Fatal(err)
		}

		env := defaultEnvironment
		if len(args) == 1 {
			env = args[0]
		}

		tag := ""
		skipBuildNPublishFlag := cmd.Flags().Lookup("skip-build-n-publish")
		skipBuildNPublish, err := strconv.ParseBool(skipBuildNPublishFlag.Value.String())
		if err != nil {
			log.Errorf("Failed reading skip-build-n-publish flag %s.", err)
			return
		}

		if skipBuildNPublish {
			log.Info("skipping build and publish procedures")
		} else {
			skipCacheFlag := cmd.Flags().Lookup("skip-cache")
			skipCache, err := strconv.ParseBool(skipCacheFlag.Value.String())
			if err != nil {
				log.Errorf("Failed reading skip-cache flag %s.", err)
				return
			}

			builder := building.NewBuilder(ctx)

			// Build current docker project & create a tag for the deploy
			tag, err = builder.Build(env, skipCache)
			if err != nil {
				log.Errorf("Failed building current docker project %s - %s.", env, err)
				return
			}

			// Publish the tag to the docker repository
			// err = builder.Publish(tag)
			// if err != nil {
			// 	log.Errorf("Failed publishing tag %s %s - %s.", tag, env, err)
			// 	return
			// }
		}

		//Deploy the tag
		d := deployment.NewDeployer(ctx)
		d.Execute(ctx, env, tag, false)
	}
}
