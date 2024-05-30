package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(filename string, value float64){
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func GetFloatFromFile(filename string) (value float64, err error){
	// balanceText := fmt.Sprint(balance)
	data, err := os.ReadFile(filename)
	if err != nil{
		return 1000, errors.New("failed to read value")
	}
	valueText := string(data)
	value, err = strconv.ParseFloat(valueText, 64)
	if err != nil{
		return 1000, errors.New("failed to parse value")
	}
	return value, nil
}