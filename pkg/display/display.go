package display

import "fmt"

func Show(input string, err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("%s\n", input)
}
