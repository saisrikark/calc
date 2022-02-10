package helpers

import "strconv"

func ArgsToFloat(args []string) ([]float64, error) {
	ret := make([]float64, len(args))
	for i, val := range args {
		f, err := strconv.ParseFloat(val, 64)
		ret[i] = f
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}
