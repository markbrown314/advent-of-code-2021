package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	uniqueNumberCount := 0
	fileInput, err := ioutil.ReadFile("day-8-input.txt")
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
			//fmt.Println(outputStr, len(outputStr))
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
