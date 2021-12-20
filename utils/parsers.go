package utils

import (
	"strconv"
	"strings"
)

func ConvStrToIntList(numList string, delimter string) ([]int, error) {
	numArray := make([]int, 0, 64)
	for _, numStr := range strings.Split(numList, delimter) {
		numInt, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numArray = append(numArray, numInt)
	}
	return numArray, nil
}
