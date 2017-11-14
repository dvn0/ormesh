// Copyright © 2017 Casey Marshall
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cmars/ormesh/config"
)

// remoteShowCmd represents the remoteShow command
var remoteShowCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		withConfig(func(cfg *config.Config) error {
			remoteName := args[0]
			if !IsValidRemoteName(remoteName) {
				return errors.Errorf("invalid remote %q", remoteName)
			}
			for _, remote := range cfg.Node.Remotes {
				if remote.Name == remoteName {
					fmt.Println(remote)
					return nil
				}
			}
			return errors.Errorf("no such remote %q", remoteName)
		})
	},
}

func init() {
	remoteCmd.AddCommand(remoteShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteShowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteShowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
