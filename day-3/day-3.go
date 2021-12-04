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

func countBits(binArray []uint32, binLength int, pos int) ([]uint32, []uint32) {
	binArrayZeroes := make([]uint32, 0, len(binArray))
	binArrayOnes := make([]uint32, 0, len(binArray))

	for j := 0; j < len(binArray); j++ {
		switch (binArray[j] >> (binLength - pos - 1)) & 1 {
		case 0:
			binArrayZeroes = append(binArrayZeroes, binArray[j])
		case 1:
			binArrayOnes = append(binArrayOnes, binArray[j])
		}
	}

	return binArrayZeroes, binArrayOnes
}

func part1(binArray []uint32, binLength int) {
	var gammaRate uint32 = 0
	var epsilonRate uint32 = 0

	for i := 0; i < binLength; i++ {
		bitSetting := 0
		binArrayZeroes, binArrayOnes := countBits(binArray, binLength, i)
		if len(binArrayOnes) > len(binArrayZeroes) {
			bitSetting = 1
		}
		gammaRate = gammaRate<<1 | uint32(bitSetting)
		epsilonRate = epsilonRate<<1 | uint32(^bitSetting&0x1)
	}

	fmt.Printf("part 1: power consumption %v\n", gammaRate*epsilonRate)

}

func part2(binArray []uint32, binLength int) {
	o2BinArray := make([]uint32, len(binArray))
	cO2BinArray := make([]uint32, len(binArray))

	copy(o2BinArray, binArray)
	copy(cO2BinArray, binArray)

	for i := 0; i < binLength; i++ {
		if len(o2BinArray) == 1 {
			break
		}
		binArrayZeroes, binArrayOnes := countBits(o2BinArray, binLength, i)
		if len(binArrayOnes) >= len(binArrayZeroes) {
			o2BinArray = make([]uint32, len(binArrayOnes))
			copy(o2BinArray, binArrayOnes)
		} else {
			o2BinArray = make([]uint32, len(binArrayZeroes))
			copy(o2BinArray, binArrayZeroes)
		}
	}

	o2ScrubberRating := o2BinArray[0]

	for i := 0; i < binLength; i++ {
		if len(cO2BinArray) == 1 {
			break
		}
		binArrayZeroes, binArrayOnes := countBits(cO2BinArray, binLength, i)
		if len(binArrayZeroes) <= len(binArrayOnes) {
			cO2BinArray = make([]uint32, len(binArrayZeroes))
			copy(cO2BinArray, binArrayZeroes)
		} else {
			cO2BinArray = make([]uint32, len(binArrayOnes))
			copy(cO2BinArray, binArrayOnes)
		}
	}

	cO2ScrubberRating := cO2BinArray[0]

	fmt.Println("part 2: life support rating:", o2ScrubberRating*cO2ScrubberRating)
}

func main() {
	input, err := ioutil.ReadFile("day-3-input-1.txt")
	if err != nil {
		log.Fatalf("error reading input file %v", err)
	}
	binArray, binLength := convertInput(input)
	part1(binArray, binLength)
	part2(binArray, binLength)
}
