package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type FloorTile struct {
	height int
	id     int
}

const HeightMax = 9

// coordinate offsets above, below, left, and right
var coordArray []Coord

func init() {
	coordArray = []Coord{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
}

// determine if the coordinate is the minimum height
func isMinHeight(c Coord, floorMap map[Coord]FloorTile) bool {
	floorTile, valid := floorMap[Coord{c.X, c.Y}]
	if !valid {
		log.Fatalf("invalid coordinate (%v,%v)\n", c.X, c.Y)
	}
	for _, coord := range coordArray {
		checkFloorTile, valid := floorMap[Coord{c.X + coord.X, c.Y + coord.Y}]
		// map out of bound tiles as maximum height tile
		if !valid {
			checkFloorTile = FloorTile{HeightMax, 0}
		}
		if floorTile.height >= checkFloorTile.height {
			return false
		}
	}
	return true
}

// mark a basin (region of connected tiles with specified id)
func checkFloorTile(c Coord, id int, floorMap map[Coord]FloorTile) int {
	tile, valid := floorMap[Coord{c.X, c.Y}]

	// reached out of bounds
	if !valid {
		return 0
	}

	// reached border of basin
	if tile.height == 9 {
		return 0
	}

	// reached marked tile in same basin
	if tile.id == id {
		return 0
	}

	// assert if reached marked tile from other basin
	if tile.id != 0 {
		log.Fatal("error: touched other basin")
	}

	// mark tile
	tile.id = id
	floorMap[c] = tile
	size := 1

	for _, coord := range coordArray {
		// recusively find adjacent floor tiles (sum size of found tiles)
		size += checkFloorTile(Coord{c.X + coord.X, c.Y + coord.Y}, id, floorMap)
	}

	return size
}

// find minimum heights and sum risk levels
func part1(floorMap map[Coord]FloorTile) {
	sumRiskLevel := 0
	for key, value := range floorMap {
		if isMinHeight(key, floorMap) {
			sumRiskLevel += value.height + 1
		}
	}

	fmt.Printf("part 1: sum of risk levels %v\n", sumRiskLevel)
}

// find three largest size basins and multiply them
// a basin is a set of connected tiles of height < 9
func part2(floorMap map[Coord]FloorTile) {
	id := 0
	basin := make([]int, 0)
	for key, value := range floorMap {
		if value.height == 9 {
			continue
		}
		// untouched basin found mark all connected tiles in basin with new id
		if value.id == 0 {
			id++
			size := checkFloorTile(key, id, floorMap)
			basin = append(basin, size)
		}
	}
	if len(basin) < 3 {
		log.Fatalln("Too few basins detected")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basin)))
	fmt.Printf("part 2: results of three largest basins %v\n", basin[0]*basin[1]*basin[2])
}

func main() {
	floorMap := make(map[Coord]FloorTile)
	fmt.Printf("Problem Day #9\n")

	// load input file and convert to map of floor tiles
	fileInput, err := ioutil.ReadFile("day-9-input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v\n", err)
	}
	inputLines := strings.Split(string(fileInput), "\n")
	for y, line := range inputLines {
		for x, input := range line {
			height, err := strconv.Atoi(string(input))
			if err != nil {
				log.Fatalf("Failed to convert string to int %v", err)
			}
			floorMap[Coord{x, y}] = FloorTile{height, 0}
		}
	}
	part1(floorMap)
	part2(floorMap)
}
