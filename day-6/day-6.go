package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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

	strData := strings.Split(string(fileInput), ",")
	for _, initTimerStr := range strData {
		timer, err := strconv.Atoi(initTimerStr)
		if err != nil {
			log.Fatalf("error: cannot convert string to int %v\n", err)
		}

		if timer < 0 || timer > MaxTimer {
			log.Fatalf("error: invalid timer value %v\n", timer)
		}

		lanternFish[timer]++
	}

	for true {
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
	fmt.Println("part1", calculateFish(256))
}
