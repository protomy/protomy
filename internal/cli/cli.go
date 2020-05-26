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

package cli

import (
	"github.com/spf13/cobra"

	"github.com/protomy/protomy/internal/cli/command"
	"github.com/protomy/protomy/internal/logging"
)

const (
	cliName        = "protomy"
	cliDescription = "Generate projects from templates in version control"
)

var globalFlags = command.GlobalFlags{}

var rootCmd = &cobra.Command{
	Use:              cliName,
	Short:            cliDescription,
	SuggestFor:       []string{"protomy"},
	PersistentPreRun: persistentPreRunFunc,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&globalFlags.ConfigFilename, "config", "C", "", "Configuration file to use")
	rootCmd.PersistentFlags().StringToStringVarP(&globalFlags.ConfigOverrides, "config-option", "c", map[string]string{}, "Configuration options to override")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Quiet, "quiet", "q", false, "Silence output")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.Quiet, "silent", false, "Alias of -q and --quiet")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Verbose, "verbose", "v", false, "Output extra debugging information")

	rootCmd.AddCommand(
		command.NewCheckCommand(),
		command.NewCompletionCommand(rootCmd),
		command.NewConfigCommand(),
		command.NewEnvCommand(),
		command.NewDownloadCommand(),
		command.NewGenerateCommand(),
		command.NewValidateCommand(),
	)
}

func Start() {
	err := rootCmd.Execute()
	if err != nil {
		command.ExitWithError(command.ExitError, err)
	}
}

func persistentPreRunFunc(cmd *cobra.Command, args []string) {
	logger := logging.New()
	command.SetLogger(logger)
}
