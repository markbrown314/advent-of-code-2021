package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func calculateFish(maxDays int) int {
	var lanternFishList []int
	day := 0
	i := 0
	fileInput, err := ioutil.ReadFile("day-6-input.txt")
	strData := strings.Split(string(fileInput), ",")
	for _, initTimerStr := range strData {
		initTimerInt, err := strconv.Atoi(initTimerStr)
		if err != nil {
			log.Fatalf("error: cannot convert string to int %v\n", err)
		}
		lanternFishList = append(lanternFishList, initTimerInt)
	}

	if err != nil {
		log.Fatalf("error: cannot open file %v\n", fileInput)
	}

	for true {
		if lanternFishList[i] == 0 {
			lanternFishList[i] = 6
			lanternFishList = append(lanternFishList, 9)
		} else {
			lanternFishList[i]--
		}
		i = (i + 1) % len(lanternFishList)
		if i == 0 {
			day += 1
			if day == maxDays {
				break
			}
		}
	}
	return len(lanternFishList)
}
func main() {
	fmt.Println("part1", calculateFish(80))
}
