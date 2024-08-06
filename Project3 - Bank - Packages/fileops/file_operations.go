package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0664)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("failed to find file")
	}

	valueText := string(data)

	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("failed to parse balance file value")
	}

	return value, nil
}
