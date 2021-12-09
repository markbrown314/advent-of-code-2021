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

func fixCoord(in PointPair) (out PointPair) {
	out = in

	if in.X1 > in.X2 {
		out.X2 = in.X1
		out.X1 = in.X2
	}

	if in.Y1 > in.Y2 {
		out.Y2 = in.Y1
		out.Y1 = in.Y2
	}

	out.deltaX = out.X2 - out.X1
	out.deltaY = out.Y2 - out.Y1

	return
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
		tmpPoint = fixCoord(tmpPoint)
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

func part1(fileName string) {
	var score int
	ventField := inputFile(fileName)

	for _, point := range ventField.points {

		// skip diagonals
		if point.deltaX > 0 && point.deltaY > 0 {
			continue
		}

		if point.deltaY == 0 {
			for x := point.X1; x <= point.X2; x++ {
				ventField.field[x][point.Y1]++
			}
		}

		if point.deltaX == 0 {
			for y := point.Y1; y <= point.Y2; y++ {
				ventField.field[point.X1][y]++
			}
		}
	}

	for x := 0; x < initX; x++ {
		for y := 0; y < initY; y++ {
			if ventField.field[x][y] > 1 {
				score++
			}
		}
	}

	fmt.Printf("part1 score %v\n", score)

}

func part2(fileName string) {
	var score int
	ventField := inputFile(fileName)

	for _, point := range ventField.points {

		if point.deltaY == 0 {
			for x := point.X1; x <= point.X2; x++ {
				ventField.field[x][point.Y1]++
			}
		}

		if point.deltaX == 0 {
			for y := point.Y1; y <= point.Y2; y++ {
				ventField.field[point.X1][y]++
			}
		}
	}

	for x := 0; x < initX; x++ {
		for y := 0; y < initY; y++ {
			if ventField.field[x][y] > 1 {
				score++
			}
		}
	}

	fmt.Printf("part2 score %v\n", score)

}

func main() {
	part1("day-5-test.txt")
}
