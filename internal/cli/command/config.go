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

// NewConfigCommand builds and returns a *cobra.Command for the "config" command.
func NewConfigCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Read or write configuration",
		Run:   configCommandFunc,
	}
}

func configCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Println("Running config command...")
}
