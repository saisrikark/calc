package format

import "fmt"

func FormatOutput(input float64, err error) (string, error) {
	return fmt.Sprintf("%v", input), nil
}
