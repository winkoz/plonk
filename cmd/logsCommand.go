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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io/log"
	"github.com/winkoz/plonk/internal/management"
)

func addLogsCommand(rootCmd *cobra.Command) {
	var logsCmd = &cobra.Command{
		Use:     "logs",
		Short:   "Gets the logs for an application deployed by plonk, filtering by the environment provided.",
		Run:     newLogsCommandHandler(),
		Example: "plonk diff",
		Args:    cobra.MaximumNArgs(1),
	}

	rootCmd.AddCommand(logsCmd)
}

func newLogsCommandHandler() CobraHandler {
	return func(cmd *cobra.Command, args []string) {
		ctx, err := config.NewContextFromFile()
		if err != nil {
			log.Fatal(err)
		}

		env := defaultEnvironment
		if len(args) == 1 {
			env = args[0]
		}

		desiredComponent := retrieveComponent(ctx, env)

		m := management.NewManager(ctx)
		m.GetLogs(env, desiredComponent)
	}
}

func retrieveComponent(ctx config.Context, env string) (desiredComponent *string) {
	// Extract components for desired env
	components := ctx.Components(env)
	componentsCount := len(components)

	// Only ask if there is more than 1 component
	if componentsCount > 1 {
		input := ""
		// Keep asking until we receive a valid option
		for repeat := true; repeat; {
			fmt.Println("Choose which component to retrieve the logs from:")
			for idx, component := range components {
				fmt.Printf("\t%d. %s\n", idx, component)
			}
			// Inject the «All» option
			fmt.Printf("\t%d. All\n", componentsCount)

			// Read the selection from the user
			_, _ = fmt.Scan(&input)
			option, err := strconv.Atoi(input)

			// Validate if the option is a number and is within the range of available components
			if repeat = (err != nil || option > componentsCount); !repeat {
				component := "All"
				if option < componentsCount { // If option is within the list of components read it
					desiredComponent = &(components[option])
					component = *desiredComponent
				}
				log.Infof("Selected logs for %s component(s).", component)
			} else {
				fmt.Println("Invalid option provided")
			}
		}
	}

	return
}
