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
	"os"

	"github.com/spf13/cobra"
)

// NewCompletionCommand builds and returns a *cobra.Command for the "completion" command.
func NewCompletionCommand(rootCmd *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:       "completion",
		Short:     "Generate shell completions",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{"bash", "fish", "powershell", "zsh"},
		Run:       completionCommandFunc(rootCmd),
	}
}

func completionCommandFunc(rootCmd *cobra.Command) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "bash":
			err = rootCmd.GenBashCompletion(os.Stdout)
		case "fish":
			err = rootCmd.GenFishCompletion(os.Stdout, false)
		case "powershell":
			err = rootCmd.GenPowerShellCompletion(os.Stdout)
		case "zsh":
			err = rootCmd.GenZshCompletion(os.Stdout)
		}

		if err != nil {
			ExitWithError(ExitError, err)
		}
	}
}
