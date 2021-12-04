package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func convertInput(input []byte) ([]uint32, int) {
	binArray := make([]uint32, 0, 1000)
	lines := strings.Split(string(input), "\n")
	binLength := 0
	for _, valueStr := range lines {
		valueInt, err := strconv.ParseUint(valueStr, 2, 32)
		if err != nil {
			log.Fatalf("error during string conversion %v\n", err)
		}
		if len(valueStr) > binLength {
			binLength = len(valueStr)
		}
		binArray = append(binArray, uint32(valueInt))
	}
	return binArray, binLength
}

func countBits(binArray []uint32, binLength int, pos int) int {
	one_count := 0
	zero_count := 0
	for j := 0; j < len(binArray); j++ {
		switch (binArray[j] >> (binLength - pos - 1)) & 1 {
		case 0:
			zero_count++
		case 1:
			one_count++
		}
	}

	if zero_count > one_count {
		return 0
	}

	return 1
}

func part1(binArray []uint32, binLength int) {

	var gammaRate uint32 = 0
	var epsilonRate uint32 = 0

	for i := 0; i < binLength; i++ {
		bitSetting := countBits(binArray, binLength, i)
		gammaRate = gammaRate<<1 | uint32(bitSetting)
		epsilonRate = epsilonRate<<1 | uint32(^bitSetting&0x1)
	}

	fmt.Printf("part 1: power consumption %v\n", gammaRate*epsilonRate)

}

func main() {
	input, err := ioutil.ReadFile("day-3-input-1.txt")
	if err != nil {
		log.Fatalf("error reading input file %v", err)
	}
	binArray, binLength := convertInput(input)
	part1(binArray, binLength)
}
