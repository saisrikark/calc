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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds numbers separated by space",
	Long:  `given a list of numbers separated by a space, sums them up and shows the result`,
	Run: func(cmd *cobra.Command, args []string) {
		floatStatus, _ := cmd.Flags().GetBool("float")
		if floatStatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "add floating numbers")
}

func addInt(args []string) {
	sum := 0
	for _, ival := range args {
		if itemp, err := strconv.Atoi(ival); err != nil {
			fmt.Println(err.Error())
			return
		} else {
			sum += itemp
		}
	}
	fmt.Printf("%d\n", sum)
}

func addFloat(args []string) {
	sum := 0.0
	for _, ival := range args {
		if itemp, err := strconv.ParseFloat(ival, 64); err != nil {
			fmt.Println(err.Error())
			return
		} else {
			sum += itemp
		}
	}
	fmt.Printf("%f\n", sum)
}
