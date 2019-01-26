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
	"encoding/json"
	"fmt"

	"github.com/iost-official/go-iost/account"
	"github.com/iost-official/go-iost/common"
	"github.com/spf13/cobra"
)

type key struct {
	Algorithm string
	Pubkey    string
	Seckey    string
}

// keyCmd represents the keyPair command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Create a key pair",
	Long:  `Create a key pair`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n, err := account.NewKeyPair(nil, sdk.GetSignAlgo())
		if err != nil {
			return fmt.Errorf("failed to new key pair: %v", err)
		}

		var k key
		k.Algorithm = n.Algorithm.String()
		k.Pubkey = common.Base58Encode(n.Pubkey)
		k.Seckey = common.Base58Encode(n.Seckey)

		ret, err := json.MarshalIndent(k, "", "    ")
		if err != nil {
			return fmt.Errorf("failed to marshal: %v", err)
		}
		fmt.Println(string(ret))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)
}
