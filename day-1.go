package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// convert array of string of ints to an array of ints
func convArrStringtoInt(strArray []string) []int {
	intArray := make([]int, 0, 64)
	for i := 0; i < len(strArray); i++ {
		intInput, err := strconv.Atoi(strArray[i])
		if err != nil {
			log.Fatalf("data conversion error for %v (%v)", strArray[i], err)
		}
		intArray = append(intArray, intInput)
	}
	return intArray
}

func tallyIncrease(intArray []int) (count int) {
	for i, element := range intArray[1:] {
		if element > intArray[i] {
			count++
		}
	}
	return
}

func sumSlidingWindow(windowSize int, intArray []int) []int {
	base := 0
	sumArray := make([]int, 0, 64)

	for i := 0; i <= len(intArray); i++ {
		if i-base < windowSize {
			continue
		}
		sum := 0
		for j := base; j < i; j++ {
			sum += intArray[j]
		}
		sumArray = append(sumArray, sum)
		base++
	}
	return sumArray
}

func main() {

	input, err := ioutil.ReadFile("day-1-input-1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	// part 1
	depthInput := convArrStringtoInt(strings.Split(string(input), "\n"))
	fmt.Printf("Part 1: depth has increased %v times\n", tallyIncrease(depthInput))

	// part 2
	sumArray := sumSlidingWindow(3, depthInput)
	fmt.Printf("Part 2: depth has increased %v times\n", tallyIncrease(sumArray))
}
