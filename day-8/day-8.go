package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var segmentMap map[int]int
var wireMap map[rune]int

func init() {

	segmentMap = make(map[int]int, 7)
	/* number = 0bgfedcba
	    aaaa
	   b    c
	   b    c
	    dddd
	   e    f
	   e    f
	    gggg
	*/

	// bits are offset by one due to 0 = map failure
	segmentMap[0] = 0b11101110
	segmentMap[1] = 0b01001000
	segmentMap[2] = 0b10111010
	segmentMap[3] = 0b11011010
	segmentMap[4] = 0b01011100
	segmentMap[5] = 0b11010110
	segmentMap[6] = 0b11110110
	segmentMap[7] = 0b01001010
	segmentMap[8] = 0b11111110
	segmentMap[9] = 0b11011110

	wireMap = make(map[rune]int)
	wireMap['a'] = 1
	wireMap['b'] = 2
	wireMap['c'] = 3
	wireMap['d'] = 4
	wireMap['e'] = 5
	wireMap['f'] = 6
	wireMap['g'] = 7
}

func main() {
	uniqueNumberCount := 0
	fileInput, err := ioutil.ReadFile("day-8-test.txt")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	re, err := regexp.Compile("[||\n]")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	inputStr := re.Split(string(fileInput), -1)
	for i := range inputStr {
		if i%2 == 0 {
			continue
		}
		for _, outputStr := range strings.Fields(inputStr[i]) {
			switch len(outputStr) {
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 7:
				uniqueNumberCount++
			}
		}
	}
	fmt.Println("part1", uniqueNumberCount)
}
