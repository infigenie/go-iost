// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package iwallet

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// stateCmd prints the state of blockchain.
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Get blockchain and node state",
	Long:  `Get blockchain and node state`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n, err := sdk.getNodeInfo()
		if err != nil {
			return fmt.Errorf("cannot get node info: %v", err)
		}
		fmt.Print(strings.TrimRight(marshalTextString(n), "}"))
		c, err := sdk.getChainInfo()
		if err != nil {
			return fmt.Errorf("cannot get chain info: %v", err)
		}
		fmt.Println(strings.Replace(marshalTextString(c), "{\n", "", 1))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stateCmd)
}
