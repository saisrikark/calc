package div

import (
	"calc/pkg/operations/helpers"
)

func DivOp(_ DivOptions, args []string) (float64, error) {
	numbers, err := helpers.ArgsToFloat(args)
	if err != nil {
		return 0, err
	}
	return simpleDiv(numbers[0], numbers[1])
}

func simpleDiv(dividend float64, divisor float64) (float64, error) {
	return dividend / divisor, nil
}
