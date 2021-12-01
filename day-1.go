package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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

func part1(depthInput []int) {
	prev := 0
	increased := 0

	for i := 0; i < len(depthInput); i++ {

		depth := depthInput[i]

		if prev == 0 {
			fmt.Printf("%v (N/A - no previous measurement)\n", depth)
			prev = depth
			continue
		}

		fmt.Print(depth, " ")

		if depth == prev {
			fmt.Print("(no change)")
		}

		if depth > prev {
			fmt.Print("(increased)")
			increased++
			prev = depth
		}

		if depth < prev {
			fmt.Print("(decreased)")
			prev = depth
		}

		fmt.Println()
	}

	fmt.Printf("depth has increased %v times", increased)
}

func part2(depthInput []int) {
	base := 0
	prev := 0
	increased := 0

	for i := 0; i <= len(depthInput); i++ {

		if i-base < 3 {
			continue
		}

		sum := 0
		for j := base; j < i; j++ {
			sum += depthInput[j]
		}

		base++

		if prev == 0 {
			fmt.Printf("%v (N/A - no previous sum)\n", sum)
			prev = sum
			continue
		}

		fmt.Print(sum, " ")

		if sum == prev {
			fmt.Print("(no change)")
		}

		if sum > prev {
			fmt.Print("(increased)")
			increased++
			prev = sum
		}

		if sum < prev {
			fmt.Print("(decreased)")
			prev = sum
		}

		fmt.Println()
	}

	fmt.Printf("depth has increased %v times", increased)
}

func main() {

	input, err := ioutil.ReadFile("day-1-input-1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	sliceInput := strings.Split(string(input), "\n")

	depthInput := convArrStringtoInt(sliceInput)

	part1(depthInput)
	part2(depthInput)
}
