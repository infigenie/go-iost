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

	"github.com/spf13/cobra"
)

// receiptCmd represents the receipt command.
var receiptCmd = &cobra.Command{
	Use:   "receipt",
	Short: "Find receipt",
	Long:  `Find receipt by transaction hash`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			fmt.Println("Please enter the transaction hash")
			return
		}
		txReceipt, err := sdk.GetTxReceiptByTxHash(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(marshalTextString(txReceipt))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(receiptCmd)
}
