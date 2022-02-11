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
	add "calc/pkg/operations/add"
	check "calc/pkg/validate"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds numbers separated by space",
	Long:  `adds numbers separated by space`,
	Run:   addOp,
}

func init() {
	rootCmd.AddCommand(addCmd)
	initAddFlags()
}

func initAddFlags() {
	addCmd.Flags().BoolP("ignoreFloat", "", false, "ignore the floating part of each number")
	addCmd.Flags().BoolP("ignoreDecimal", "", false, "ignore the decimal part of each number")
}

func addOp(cmd *cobra.Command, args []string) {
	var opts add.AddOptions
	var err error
	var res string
	var fRes float64

	if ok, err := check.IsValid(check.Add, args); !ok || err != nil {
		display.Show("", err)
		return
	}
	if opts, err = fetchAddOptions(cmd); err != nil {
		display.Show("", err)
		return
	}
	if fRes, err = add.AddOp(opts, args); err != nil {
		display.Show("", err)
		return
	}
	if res, err = format.FormatOutput(fRes, err); err != nil {
		display.Show(res, err)
		return
	}

	display.Show(res, nil)
}

func fetchAddOptions(cmd *cobra.Command) (add.AddOptions, error) {
	ignoreFloat, _ := cmd.Flags().GetBool("ignoreFloat")
	ignoreDecimal, _ := cmd.Flags().GetBool("ignoreDecimal")
	addOpts := add.AddOptions{
		IgnoreFloat:   ignoreFloat,
		IgnoreDecimal: ignoreDecimal,
	}
	return addOpts, nil
}
