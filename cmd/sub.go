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
	sub "calc/pkg/operations/sub"
	check "calc/pkg/validate"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subtract two numbers",
	Long:  `Subtract two numbers`,
	Run:   subOp,
}

func init() {
	rootCmd.AddCommand(subCmd)
}

func subOp(cmd *cobra.Command, args []string) {
	var opts sub.SubOptions
	var err error
	var res string
	var fRes float64

	if ok, err := check.IsValid(check.Sub, args); !ok || err != nil {
		display.Show("", err)
		return
	}
	if opts, err = fetchSubOptions(cmd); err != nil {
		display.Show("", err)
		return
	}
	if fRes, err = sub.SubOp(opts, args); err != nil {
		display.Show("", err)
		return
	}
	if res, err = format.FormatOutput(fRes, err); err != nil {
		display.Show(res, err)
		return
	}

	display.Show(res, nil)
}

func fetchSubOptions(cmd *cobra.Command) (sub.SubOptions, error) {
	return sub.SubOptions{}, nil
}
