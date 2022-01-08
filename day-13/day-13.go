package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func stringToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("error could not convert %v to int (%v)", str, err)
	}
	return value
}

func foldAlong(c int, pivot int) int {
	a := c - pivot
	return pivot - a
}

func foldDotMap(coordMap map[Coord]bool, pivot int, along_x bool) {
	for coord := range coordMap {

		c := coord.Y
		if along_x {
			c = coord.X
		}

		//fmt.Printf("coord %v x?%v %v\n", coord, along_x, pivot)

		if c <= pivot {
			continue
		}

		c = foldAlong(c, pivot)
		var new_coord Coord

		if along_x {
			new_coord = Coord{X: c, Y: coord.Y}
		} else {
			new_coord = Coord{X: coord.X, Y: c}
		}

		//fmt.Printf("moved here %v\n", new_coord)
		delete(coordMap, coord)
		coordMap[new_coord] = true
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
	fileInput, err := ioutil.ReadFile("day-13-input.txt")
	foldCount := 0
	maxX := -1
	maxY := -1
	if err != nil {
		log.Fatalf("Error loading data file %v\n", err)
	}

	coord_re := regexp.MustCompile("^[0-9]*,[0-9]*$")
	fold_re := regexp.MustCompile("^fold along [x|y]=")
	fold_along_x_re := regexp.MustCompile("^fold along x=")

	dotMap := make(map[Coord]bool)
	for _, line := range strings.Split(string(fileInput), "\n") {
		if coord_re.MatchString(line) {
			c := strings.Split(line, ",")
			x := stringToInt(c[0])
			y := stringToInt(c[1])
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			dotMap[Coord{X: x, Y: y}] = true
			continue
		}

		if fold_re.MatchString(line) {
			pivot := strings.Split(line, "=")
			along_x := fold_along_x_re.MatchString(line)
			//fmt.Printf("Fold here x?%v pivot %v\n", along_x, pivot[1])
			foldDotMap(dotMap, stringToInt(pivot[1]), along_x)
			foldCount++
			if along_x {
				maxX /= 2
			} else {
				maxY /= 2
			}
		}

		if foldCount == 1 {
			fmt.Printf("part 1: visible coordinates after fold %v\n", len(dotMap))
		}
	}
	printDots(dotMap, maxX, maxY)
}
