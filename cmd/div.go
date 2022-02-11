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
	div "calc/pkg/operations/div"
	check "calc/pkg/validate"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "division opeartion on two numbers",
	Long:  `division opeartion on two numbers`,
	Run:   divOp,
}

func init() {
	rootCmd.AddCommand(divCmd)
}

func divOp(cmd *cobra.Command, args []string) {
	var opts div.DivOptions
	var err error
	var res string
	var fRes float64

	if ok, err := check.IsValid(check.Div, args); !ok || err != nil {
		display.Show("", err)
		return
	}
	if opts, err = fetchDivOptions(cmd); err != nil {
		display.Show("", err)
		return
	}
	if fRes, err = div.DivOp(opts, args); err != nil {
		display.Show("", err)
		return
	}
	if res, err = format.FormatOutput(fRes, err); err != nil {
		display.Show(res, err)
		return
	}

	display.Show(res, nil)
}

func fetchDivOptions(cmd *cobra.Command) (div.DivOptions, error) {
	return div.DivOptions{}, nil
}
