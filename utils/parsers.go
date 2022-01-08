package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ConvStrToIntList(numList string, regex string) ([]int, error) {
	numArray := make([]int, 0, 64)
	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	numList = strings.TrimSpace(numList)
	for _, numStr := range re.Split(numList, -1) {
		numInt, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numArray = append(numArray, numInt)
	}
	return numArray, nil
}

func StringToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("error could not convert %v to int (%v)", str, err)
	}
	return value
}
