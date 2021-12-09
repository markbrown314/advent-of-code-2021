package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// TODO should be dynamic
const (
	initX = 1000
	initY = 1000
)

type PointPair struct {
	X1     int
	Y1     int
	X2     int
	Y2     int
	deltaX int
	deltaY int
}

type hydrothermalField struct {
	field  [][]int
	points []PointPair
}

func inputFile(fileName string) (ventField hydrothermalField) {
	fileInput, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalf("error: cannot open file %v\n", fileInput)
	}

	ventField.field = make([][]int, initY, initY)
	ventField.points = make([]PointPair, 0, 64)

	for i := range ventField.field {
		ventField.field[i] = make([]int, initX, initX)
	}

	inputLines := strings.Split(string(fileInput), "\n")

	for _, input := range inputLines {
		inputPair := strings.Split(string(input), " -> ")
		tmpPoint := PointPair{}
		tmpPoint.X1, tmpPoint.Y1 = parseCoord(inputPair[0])
		tmpPoint.X2, tmpPoint.Y2 = parseCoord(inputPair[1])
		tmpPoint.deltaX = tmpPoint.X2 - tmpPoint.X1
		tmpPoint.deltaY = tmpPoint.Y2 - tmpPoint.Y1
		ventField.points = append(ventField.points, tmpPoint)
	}
	return ventField
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

func scanVentField(fileName string, skipDiagonals bool) int {
	var score int

	ventField := inputFile(fileName)

	for _, point := range ventField.points {
		dirX := 0
		dirY := 0

		// skip diagonals
		if point.deltaX != 0 && point.deltaY != 0 && skipDiagonals {
			continue
		}

		if point.deltaX > 0 {
			dirX = 1
		}

		if point.deltaX < 0 {
			dirX = -1
		}

		if point.deltaY > 0 {
			dirY = 1
		}

		if point.deltaY < 0 {
			dirY = -1
		}
		x := point.X1
		y := point.Y1

		for true {
			ventField.field[x][y]++
			if x == point.X2 && y == point.Y2 {
				break
			}
			x += dirX
			y += dirY
		}
	}

	// compute score
	for x := 0; x < initX; x++ {
		for y := 0; y < initY; y++ {
			if ventField.field[x][y] > 1 {
				score++
			}
		}
	}

	return score
}

func main() {
	fileName := "day-5-input.txt"
	fmt.Println("part1", scanVentField(fileName, true))
	fmt.Println("part2", scanVentField(fileName, false))
}
