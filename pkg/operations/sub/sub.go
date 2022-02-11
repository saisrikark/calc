package sub

import "calc/pkg/operations/helpers"

func SubOp(_ SubOptions, args []string) (float64, error) {
	numbers, err := helpers.ArgsToFloat(args)
	if err != nil {
		return 0, err
	}
	return simpleSub(numbers[0], numbers[1])
}

func simpleSub(num1 float64, num2 float64) (float64, error) {
	return num1 - num2, nil
}
