package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/markbrown314/advent-of-code-2021/utils"
)

type Coord struct {
	X int
	Y int
}

// fold coordinate along pivot point
func foldAlong(c int, pivot int) int {
	a := c - pivot
	return pivot - a
}

func foldDotMap(dotMap map[Coord]bool, pivot int, alongX bool) {
	for coord := range dotMap {
		var new_coord Coord

		c := coord.Y
		if alongX {
			c = coord.X
		}

		if c <= pivot {
			continue
		}

		c = foldAlong(c, pivot)

		if alongX {
			new_coord = Coord{X: c, Y: coord.Y}
		} else {
			new_coord = Coord{X: coord.X, Y: c}
		}

		delete(dotMap, coord)
		dotMap[new_coord] = true
	}
}

func printDots(dotMap map[Coord]bool, maxX int, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if dotMap[Coord{X: x, Y: y}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Problem Day #13")

	// parse input file
	fileInput, err := ioutil.ReadFile("day-13-input.txt")
	foldCount := 0
	maxX := -1
	maxY := -1
	if err != nil {
		log.Fatalf("Error loading data file %v\n", err)
	}

	coordRe := regexp.MustCompile("^[0-9]*,[0-9]*$")
	foldRe := regexp.MustCompile("^fold along [x|y]=")
	foldAlongXRe := regexp.MustCompile("^fold along x=")

	dotMap := make(map[Coord]bool)
	for _, line := range strings.Split(string(fileInput), "\n") {
		// save location in dot map
		if coordRe.MatchString(line) {
			c := strings.Split(line, ",")
			x := utils.StringToInt(c[0])
			y := utils.StringToInt(c[1])
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			dotMap[Coord{X: x, Y: y}] = true
			continue
		}

		// fold along x or y axis
		if foldRe.MatchString(line) {
			pivot := strings.Split(line, "=")
			alongX := foldAlongXRe.MatchString(line)

			foldDotMap(dotMap, utils.StringToInt(pivot[1]), alongX)
			foldCount++

			// re adjust max
			if alongX {
				maxX /= 2
			} else {
				maxY /= 2
			}
		}

		if foldCount == 1 {
			fmt.Printf("part 1: visible coordinates after fold %v\n", len(dotMap))
		}
	}

	fmt.Println("part 2: text output:")
	printDots(dotMap, maxX, maxY)
}
