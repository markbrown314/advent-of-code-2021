package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	initX = 1000
	initY = 1000
)

type PointPair struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func checkCoord(in PointPair) (out PointPair, skip bool, deltaX int, deltaY int) {
	out = in

	if in.X1 != in.X2 && in.Y1 != in.Y2 {
		skip = true
		return
	}
	if in.X1 > in.X2 {
		out.X2 = in.X1
		out.X1 = in.X2
	}

	if in.Y1 > in.Y2 {
		out.Y2 = in.Y1
		out.Y1 = in.Y2
	}

	deltaX = out.X2 - out.X1
	deltaY = out.Y2 - out.Y1

	return
}

func parseCoord(coord string) (int, int) {
	points := strings.Split(coord, ",")
	coordInts := make([]int, 0, 2)
	for _, pointString := range points {
		pointInt, err := strconv.Atoi(pointString)
		if err != nil {
			log.Fatalf("error: cannot convert string to int %v\n", err)
		}
		coordInts = append(coordInts, pointInt)
	}
	return coordInts[0], coordInts[1]
}

func main() {
	var score int
	fileInput, err := ioutil.ReadFile("day-5-input.txt")

	if err != nil {
		log.Fatalf("error: cannot open file %v\n", fileInput)
	}

	field := make([][]int, initX, initY)
	for i := range field {
		field[i] = make([]int, initX, initY)
	}

	inputLines := strings.Split(string(fileInput), "\n")

	for _, input := range inputLines {
		inputPair := strings.Split(string(input), " -> ")
		point := PointPair{}
		point.X1, point.Y1 = parseCoord(inputPair[0])
		point.X2, point.Y2 = parseCoord(inputPair[1])
		point, skip, deltaX, deltaY := checkCoord(point)
		if skip {
			continue
		}

		if deltaY == 0 {
			for x := point.X1; x <= point.X2; x++ {
				field[x][point.Y1]++
			}
		}

		if deltaX == 0 {
			for y := point.Y1; y <= point.Y2; y++ {
				field[point.X1][y]++
			}
		}
	}

	for x := 0; x < initX; x++ {
		for y := 0; y < initY; y++ {
			if field[x][y] > 1 {
				score++
			}
		}
	}

	fmt.Printf("part1 score %v\n", score)
}
