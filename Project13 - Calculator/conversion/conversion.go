package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for index, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("error converting string to float64")
		}

		floats[index] = floatVal
	}

	return floats, nil
}
