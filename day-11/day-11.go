package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

// coordinate offsets above, below, left, right, diagonals
var coordArray []Coord

func init() {
	coordArray = []Coord{{0, -1}, {-1, 0}, {1, 0}, {0, 1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
}
func printOctopusMap(octopusMap map[Coord]int) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			power := octopusMap[Coord{x, y}]
			if power > 9 {
				fmt.Printf("*")
			} else {
				fmt.Printf("%d", power)
			}
		}
		fmt.Printf("\n")
	}
}

func updateOctopus(c Coord, octopusMap map[Coord]int, resetMap map[Coord]bool) {

	power, valid := octopusMap[c]
	if !valid {
		return
	}

	power++
	octopusMap[c] = power

	if power == 10 {
		resetMap[c] = true
		for _, coord := range coordArray {
			// recusively find adjacent floor tiles (sum size of found tiles)
			updateOctopus(Coord{c.X + coord.X, c.Y + coord.Y}, octopusMap, resetMap)
		}
	}
}

func updateOctopusMap(octopusMap map[Coord]int) (int, int) {
	flashes := 0
	resetMap := make(map[Coord]bool)
	for c := range octopusMap {
		updateOctopus(c, octopusMap, resetMap)
	}
	for c := range resetMap {
		octopusMap[c] = 0
		flashes++
	}
	return flashes, len(resetMap)
}

func main() {
	fmt.Println("Problem Day #11")
	octopusMap := make(map[Coord]int)
	// load input file
	fileInput, err := ioutil.ReadFile("day-11-input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v\n", err)
	}
	inputLines := strings.Split(string(fileInput), "\n")
	for y, line := range inputLines {
		for x, input := range line {
			power, err := strconv.Atoi(string(input))
			if err != nil {
				log.Fatalf("Failed to convert string to int %v", err)
			}
			octopusMap[Coord{x, y}] = power
		}
	}
	var totalFlashes int
	var i int
	allFlashedStep := -1
	for {
		printOctopusMap(octopusMap)
		i++
		fmt.Printf("Step %d\n", i)
		flashes, resets := updateOctopusMap(octopusMap)
		if i <= 100 {
			totalFlashes += flashes
		}

		if resets == 100 {
			allFlashedStep = i
			break
		}
		fmt.Printf("\n")
		if i >= 100 && allFlashedStep >= 0 {
			break
		}
	}
	fmt.Printf("part 1 total flashes %d\n", totalFlashes)
	fmt.Printf("part 2 all flashed step %d\n", allFlashedStep)
}
