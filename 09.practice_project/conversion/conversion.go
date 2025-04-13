package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for stringIndex, stringVal := range strings {
		floatPrise, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		floats[stringIndex] = floatPrise
	}

	return floats, nil
}
