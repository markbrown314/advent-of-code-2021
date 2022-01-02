package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var closingMap map[rune]rune
var openingMap map[rune]rune

func init() {
	closingMap = make(map[rune]rune)
	closingMap[')'] = '('
	closingMap['}'] = '{'
	closingMap[']'] = '['
	closingMap['>'] = '<'

	openingMap = make(map[rune]rune)
	openingMap['('] = ')'
	openingMap['{'] = '}'
	openingMap['['] = ']'
	openingMap['<'] = '>'
}

func main() {
	var check int
	var score int
	autoScores := make([]int, 0)
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
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				check = chunk[chr] + 1
				last = append(last, chr)
				chunk[chr] = check
				fmt.Printf("%c", chr)
			case ')':
				fallthrough
			case ']':
				fallthrough
			case '}':
				fallthrough
			case '>':
				closing := closingMap[chr]
				check = chunk[closing] - 1
				if last[len(last)-1:][0] != closing || check < 0 {
					fmt.Printf(" !%c! ", chr)
					illegal = true
				} else {
					chunk[closing] = check
					last = last[:len(last)-1]
					fmt.Printf("%c", chr)
				}
			default:
				log.Fatalf("invalid character found %c\n", chr)
			}

			if illegal {
				fmt.Printf("\n")
				fmt.Printf("** invalid %c found", chr)
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
		var autoScore int
		if len(last) != 0 && !illegal {
			fmt.Printf("** incomplete line %s\n", string(last))
			reverse := make([]rune, 0)
			for i := len(last) - 1; i >= 0; i-- {
				reverse = append(reverse, openingMap[last[i]])
			}
			for _, chr := range reverse {
				autoScore *= 5
				switch chr {
				case ')':
					autoScore += 1
				case ']':
					autoScore += 2
				case '}':
					autoScore += 3
				case '>':
					autoScore += 4
				}
			}
			fmt.Printf("reverse string %s\n", string(reverse))
			fmt.Printf("autocomplete score %v\n", autoScore)
			autoScores = append(autoScores, autoScore)
		}
	}
	sort.Ints(autoScores)
	finalAutoScore := autoScores[len(autoScores)/2]
	fmt.Printf("part 1: final score is %v\n", score)
	fmt.Printf("part 2: final score is %v\n", finalAutoScore)
}
