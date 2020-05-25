/*
 * Copyright 2020 Anthony Burns
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

type generateCommandFlags struct {
	DryRun    bool
	Templates []string
}

var (
	generateFlags = generateCommandFlags{}
)

func NewGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a project from one or more templates",
		Run:   generateCommandFunc,
	}
	cmd.Flags().BoolVarP(&generateFlags.DryRun, "dry-run", "n", false, "Print actions to be taken but do not perform them")
	cmd.Flags().StringArrayVarP(&generateFlags.Templates, "template", "t", []string{}, "Templates to use")

	return cmd
}

func generateCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Println("Running generate command...")

	if generateFlags.DryRun {
		fmt.Println("In dry-run mode")
	}

	if len(generateFlags.Templates) == 0 {
		fmt.Println("No templates specified")
	} else {
		fmt.Println("Templates (in order):")
		for i, template := range generateFlags.Templates {
			fmt.Println("  ", i+1, "-", template)
		}
	}
}
