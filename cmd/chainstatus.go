// Copyright © 2023 Weald Technology Trading.
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

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/go-string2eth"
)

// chainStatusCmd represents the chain status command.
var chainStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Obtain current status of the chain",
	Long: `Obtain current status of the chain.  For example:

    ethereal chain status

In quiet mode this will return 0 if the chain is available, otherwise 1.`,
	Run: func(_ *cobra.Command, _ []string) {
		ctx, cancel := localContext()
		defer cancel()

		baseFee, err := c.CurrentBaseFee(ctx)
		cli.ErrCheck(err, quiet, "failed to obtain current base fee")

		fmt.Printf("Base fee %v\n", string2eth.WeiToString(baseFee, true))
	},
}

func init() {
	chainCmd.AddCommand(chainStatusCmd)
	chainFlags(chainStatusCmd)
}
