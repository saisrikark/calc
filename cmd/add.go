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
	"calc/pkg/display"
	"calc/pkg/format"
	"calc/pkg/operations/add"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds numbers separated by space",
	Long:  `given a list of numbers separated by a space, sums them up and shows the result`,
	Run:   addOp,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("ignoreFloat", "", false, "ignore the floating part of each number")
	addCmd.Flags().BoolP("ignoreDecimal", "", false, "ignore the decimal part of each number")
}

func addOp(cmd *cobra.Command, args []string) {
	var opts add.AddOptions
	var err error
	var res string
	var fRes float64

	if opts, err = fetchOptions(cmd); err != nil {
		display.Show("", err)
	}
	if fRes, err = add.AddOp(opts, args); err != nil {
		display.Show("", err)
	}
	if res, err = format.FormatOutput(fRes, err); err != nil {
		display.Show(res, err)
	}

	display.Show(res, nil)
}

func fetchOptions(cmd *cobra.Command) (add.AddOptions, error) {
	ignoreFloat, _ := cmd.Flags().GetBool("ignoreFloat")
	ignoreDecimal, _ := cmd.Flags().GetBool("ignoreDecimal")
	addOpts := add.AddOptions{
		IgnoreFloat:   ignoreFloat,
		IgnoreDecimal: ignoreDecimal,
	}
	return addOpts, nil
}
