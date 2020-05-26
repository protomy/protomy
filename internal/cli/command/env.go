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

func NewEnvCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "env",
		Short: "Print information about the current environment",
		Run:   envCommandFunc,
	}
}

// NewEnvCommand builds and returns a *cobra.Command for the "env" command.
func envCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Println("Configuration overrides:")
	configOverrides, err := cmd.Flags().GetStringToString("config-option")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	for key, value := range configOverrides {
		fmt.Println("  ", key, ":", value)
	}
}
