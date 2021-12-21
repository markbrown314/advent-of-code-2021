package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/markbrown314/advent-of-code-2021/utils"
)

const (
	ResetTimer = 6
	MaxTimer   = 8
)

func calculateFish(maxDays int) uint64 {
	var lanternFish [MaxTimer + 1]uint64
	day := 0
	fileInput, err := ioutil.ReadFile("day-6-input.txt")
	if err != nil {
		log.Fatalf("error: cannot open file %v\n", fileInput)
	}

	timerList, err := utils.ConvStrToIntList(string(fileInput), ",")

	if err != nil {
		log.Fatalf("error: cannot convert string to int %v\n", err)
	}

	for _, timer := range timerList {
		if timer < 0 || timer > MaxTimer {
			log.Fatalf("error: invalid timer value %v\n", timer)
		}

		lanternFish[timer]++
	}

	for {
		spawn := lanternFish[0]
		for i := 1; i <= MaxTimer; i++ {
			lanternFish[i-1] = lanternFish[i]
		}
		lanternFish[MaxTimer] = spawn
		lanternFish[ResetTimer] += spawn

		day += 1
		if day == maxDays {
			break
		}
	}

	totalFish := uint64(0)
	for i := range lanternFish {
		totalFish += lanternFish[i]
	}
	return totalFish
}

func main() {
	fmt.Println("part1", calculateFish(80))
	fmt.Println("part2", calculateFish(256))
}
