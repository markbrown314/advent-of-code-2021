package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	input, err := ioutil.ReadFile("day-3-input-1.txt")
	if err != nil {
		log.Fatalf("error reading input file %v", err)
	}

	lines := strings.Split(string(input), "\n")

	var gamma_rate uint32 = 0
	var epsilon_rate uint32 = 0

	for i := 0; i < len(lines[i]); i++ {
		one_count := 0
		zero_count := 0
		for j := 0; j < len(lines); j++ {
			switch lines[j][i] {
			case '0':
				zero_count++
			case '1':
				one_count++
			}
		}

		var bit_setting uint32 = 1
		if zero_count > one_count {
			bit_setting = 0
		}

		if gamma_rate == 0 {
			gamma_rate |= bit_setting
		} else {
			gamma_rate = gamma_rate<<1 | bit_setting
		}

		if epsilon_rate == 0 {
			epsilon_rate |= (^bit_setting & 0x1)
		} else {
			epsilon_rate = epsilon_rate<<1 | (^bit_setting & 0x1)
		}

		//fmt.Printf("pos %v zero_count %v one_count %v gamma_rate %b epsilon_rate %b\n", i, zero_count, one_count, gamma_rate, epsilon_rate)
	}

	fmt.Printf("part 1: power consumption %v\n", gamma_rate*epsilon_rate)
}
