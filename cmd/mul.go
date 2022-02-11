/*
Copyright © 2022 Sai Srikar K

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	display "calc/pkg/display"
	format "calc/pkg/format"
	mul "calc/pkg/operations/mul"
	check "calc/pkg/validate"

	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "multiply numbers separated by a space",
	Long:  `multiply numbers separated by a space`,
	Run:   mulOp,
}

func init() {
	rootCmd.AddCommand(mulCmd)
}

func mulOp(cmd *cobra.Command, args []string) {
	var opts mul.MulOptions
	var err error
	var res string
	var fRes float64

	if ok, err := check.IsValid(check.Mul, args); !ok || err != nil {
		display.Show("", err)
		return
	}
	if opts, err = fetchMulOptions(cmd); err != nil {
		display.Show("", err)
		return
	}
	if fRes, err = mul.MulOp(opts, args); err != nil {
		display.Show("", err)
		return
	}
	if res, err = format.FormatOutput(fRes, err); err != nil {
		display.Show(res, err)
		return
	}

	display.Show(res, nil)
}

func fetchMulOptions(cmd *cobra.Command) (mul.MulOptions, error) {
	return mul.MulOptions{}, nil
}
