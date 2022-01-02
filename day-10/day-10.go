package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var check int
	var score uint32
	fmt.Println("Problem Day #10")
	// load input file
	fileInput, err := ioutil.ReadFile("day-10-input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v\n", err)
	}
	inputLines := strings.Split(string(fileInput), "\n")
	for _, line := range inputLines {
		chunk := make(map[rune]int)
		last := make([]rune, 0)
		illegal := false
		for _, chr := range line {
			switch chr {
			case '(':
				check = chunk['('] + 1
				last = append(last, chr)
				chunk['('] = check
				fmt.Printf("%c", chr)
			case '[':
				check = chunk['['] + 1
				last = append(last, chr)
				chunk['['] = check
				fmt.Printf("%c", chr)
			case '{':
				check = chunk['{'] + 1
				last = append(last, chr)
				chunk['{'] = check
				fmt.Printf("%c", chr)
			case '<':
				check = chunk['<'] + 1
				last = append(last, chr)
				chunk['<'] = check
				fmt.Printf("%c", chr)
			case ')':
				check = chunk['('] - 1
				if last[len(last)-1:][0] != '(' || check < 0 {
					fmt.Printf(" !%c! ", chr)
					illegal = true
				} else {
					chunk['('] = check
					last = last[:len(last)-1]
					fmt.Printf("%c", chr)
				}
			case ']':
				check = chunk['['] - 1
				if last[len(last)-1:][0] != '[' || check < 0 {
					fmt.Printf(" !%c! ", chr)
					illegal = true
				} else {
					chunk['['] = check
					last = last[:len(last)-1]
					fmt.Printf("%c", chr)
				}
			case '}':
				check = chunk['{'] - 1
				if last[len(last)-1:][0] != '{' || check < 0 {
					fmt.Printf(" !%c! ", chr)
					illegal = true
				} else {
					chunk['{'] = check
					last = last[:len(last)-1]
					fmt.Printf("%c", chr)
				}
			case '>':
				check = chunk['<'] - 1
				if last[len(last)-1:][0] != '<' || check < 0 {
					fmt.Printf(" !%c! ", chr)
					illegal = true
				} else {
					chunk['<'] = check
					last = last[:len(last)-1]
					fmt.Printf("%c", chr)
				}
			}
			if illegal {
				fmt.Printf("invalid %c found\n", chr)
				switch chr {
				case '}':
					score += 1197
				case ')':
					score += 3
				case ']':
					score += 57
				case '>':
					score += 25137
				}
				break
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("part 1: final score is %v\n", score)
}
