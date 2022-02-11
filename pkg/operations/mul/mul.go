package mul

import "calc/pkg/operations/helpers"

func MulOp(_ MulOptions, args []string) (float64, error) {
	numbers, err := helpers.ArgsToFloat(args)
	if err != nil {
		return 0, err
	}
	return simpleMul(numbers)
}

func simpleMul(numbers []float64) (float64, error) {
	res := 1.0
	for _, val := range numbers {
		res *= val
	}
	return res, nil
}
