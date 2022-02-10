package add

import (
	"calc/pkg/operations/helpers"
)

func AddOp(opts AddOptions, args []string) (float64, error) {
	numbers, err := helpers.ArgsToFloat(args)
	if err != nil {
		return 0, err
	}
	if opts.IgnoreDecimal {
		return addOnlyFloatingPart(numbers)
	}
	if opts.IgnoreFloat {
		return addOnlyDecimalPart(numbers)
	}
	return simpleAdd(numbers)
}

func simpleAdd(args []float64) (float64, error) {
	sum := 0.0
	for _, val := range args {
		sum += val
	}
	return sum, nil
}

func addOnlyDecimalPart(args []float64) (float64, error) {
	sum := 0.0
	for _, val := range args {
		fPart := val - float64(int64(val))
		sum += fPart
	}
	return sum, nil
}

func addOnlyFloatingPart(args []float64) (float64, error) {
	sum := 0.0
	for _, val := range args {
		sum += val - float64(int64(val))
	}
	return sum, nil
}
