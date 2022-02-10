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

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "add only even numbers from a list",
	Long:  `given a list of numbers separated by speaces, adds them`,
	Run: func(cmd *cobra.Command, args []string) {
		addEven(args)
	},
}

func init() {
	addCmd.AddCommand(evenCmd)
}

func addEven(args []string) {
	sum := 0.0
	for _, ival := range args {
		if tempVal, err := strconv.ParseFloat(ival, 64); err != nil {
			fmt.Println(err.Error())
			return
		} else {
			sum += tempVal
		}
	}
	fmt.Printf("%f\n", sum)
}
