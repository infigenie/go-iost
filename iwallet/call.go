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

	"github.com/iost-official/go-iost/ilog"
	"github.com/iost-official/go-iost/rpc/pb"
	"github.com/spf13/cobra"
)

// callCmd represents the call command that call a contract with given actions.
var callCmd = &cobra.Command{
	Use:   "call [ACTION]...",
	Short: "Call the method in contracts",
	Long: `Call the method in contracts
	Would accept arguments as call actions or load transaction request directly from proto file.
	An ACTION is a group of 3 arguments: contract name, function name, method parameters.
	The method parameters should be a string with format '["arg0","arg1",...]'.`,
	Example: `  iwallet call "token.iost" "transfer" '["iost","user0001","user0002","123.45",""]'
  iwallet call --tx_file tx.proto`,
	RunE: func(cmd *cobra.Command, args []string) error {
		trx := &rpcpb.TransactionRequest{}
		if txFile != "" {
			if len(args) != 0 {
				ilog.Warnf("load tx from file %v, will ignore cmd args %v", txFile, args)
			}
			err := loadProto(txFile, trx)
			if err != nil {
				return err
			}
		} else {
			var actions []*rpcpb.Action
			actions, err := actionsFromFlags(args)
			if err != nil {
				return err
			}
			trx, err = sdk.createTx(actions)
			if err != nil {
				return err
			}
		}
		err := sdk.LoadAccount()
		if err != nil {
			return fmt.Errorf("failed to load account: %v", err)
		}
		_, err = sdk.SendTx(trx)
		return err
	},
}

func init() {
	rootCmd.AddCommand(callCmd)
	callCmd.Flags().StringSliceVarP(&sdk.signKeys, "sign_keys", "", []string{}, "optional private key files used for signing, split by comma")
	callCmd.Flags().StringSliceVarP(&sdk.withSigns, "with_signs", "", []string{}, "optional signatures, split by comma")
	callCmd.Flags().StringVarP(&txFile, "tx_file", "", "", "load tx from this file")
}
