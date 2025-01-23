package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err :=os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("the file does not exist")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("cannot parse the value to float64")
	}

	return value, nil
}

func WriteValueToFile(value float64, fileName string){
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}