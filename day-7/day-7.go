package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/markbrown314/advent-of-code-2021/utils"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func crabRate(start int, end int) int {
	score := 0
	fuelCost := 0
	if start > end {
		temp := end
		end = start
		start = temp
	}

	for i := start; i < end; i++ {
		fuelCost++
		score += fuelCost
	}
	return score
}
func lowestIndex(data []int) int {
	lowest_pos := 0
	for i := range data {
		if data[i] < data[lowest_pos] {
			lowest_pos = i
		}
	}
	return lowest_pos
}

func main() {
	fileInput, err := ioutil.ReadFile("day-7-input.txt")
	if err != nil {
		log.Fatalf("error cannot read file %v\n", err)
	}
	pos, err := utils.ConvStrToIntList(string(fileInput), ",")
	if err != nil {
		log.Fatalf("error failure during conversion process %v\n", err)
	}

	sort.Ints(pos)
	maxPos := pos[len(pos)-1]

	score := make([]int, len(pos))

	for i := range pos {
		for j := range pos {
			score[i] = abs(pos[j]-pos[i]) + score[i]
		}
	}
	lowestPos := lowestIndex(score)
	fmt.Printf("part1 pos: %v min fuel: %v\n", pos[lowestPos], score[lowestPos])

	score2 := make([]int, maxPos)

	for i := 0; i < maxPos; i++ {
		for j := range pos {
			score2[i] += crabRate(i, pos[j])
		}
	}

	lowestPos = lowestIndex(score2)

	fmt.Printf("part2 pos: %v min fuel: %v\n", lowestPos, score2[lowestPos])
}
