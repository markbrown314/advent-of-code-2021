package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	MatrixDimension int = 5
)

func parseInputFile(fileName string) (bingoPicks []int, bingoBoard [][][]int, matchMatrix [][][]bool) {
	fileInput, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalf("error cannot read file %v\n", err)
	}

	inputLines := strings.Fields(string(fileInput))
	for _, element := range strings.Split(inputLines[0], ",") {
		pick, err := strconv.Atoi(element)
		if err != nil {
			log.Fatalf("failed to convert string to int %v\n", err)
		}
		bingoPicks = append(bingoPicks, pick)
	}

	bingoBoardInput := inputLines[1:]
	bingoBoardCount := len(inputLines[1:]) / (MatrixDimension * MatrixDimension)

	bingoBoard = make([][][]int, bingoBoardCount)
	for i := range bingoBoard {
		bingoBoard[i] = make([][]int, MatrixDimension)
		for j := range bingoBoard[i] {
			bingoBoard[i][j] = make([]int, MatrixDimension)
			pos := i*(MatrixDimension*MatrixDimension) + j*(MatrixDimension)
			for k, element := range bingoBoardInput[pos : pos+5] {
				bingoBoard[i][j][k], err = strconv.Atoi(element)
				if err != nil {
					log.Fatalf("failed to convert string to int %v\n", err)
				}
			}
		}
	}

	matchMatrix = make([][][]bool, bingoBoardCount)
	for i := range matchMatrix {
		matchMatrix[i] = make([][]bool, MatrixDimension)
		for j := range matchMatrix[i] {
			matchMatrix[i][j] = make([]bool, MatrixDimension)
		}
	}

	return
}

func checkBingoBoardVertical(matchMatrix [][]bool) (win bool) {

	for y := 0; y < MatrixDimension; y++ {
		win = true
		for x := 0; x < MatrixDimension; x++ {
			if !matchMatrix[x][y] {
				win = false
				break
			}
		}
		if win {
			return
		}
	}
	return
}

func checkBingoBoardHorizontal(matchMatrix [][]bool) (win bool) {

	for x := 0; x < MatrixDimension; x++ {
		win = true
		for y := 0; y < MatrixDimension; y++ {
			if !matchMatrix[x][y] {
				win = false
				break
			}
		}
		if win {
			return
		}
	}
	return
}

func isNumberonBoard(board [][]int, number int) (bool, int, int) {
	for x := 0; x < MatrixDimension; x++ {
		for y := 0; y < MatrixDimension; y++ {
			if board[y][x] == number {
				return true, y, x
			}
		}
	}
	return false, -1, -1
}

func computeScore(board [][]int, matchMatrix [][]bool, number int) int {
	count := 0
	for x := 0; x < MatrixDimension; x++ {
		for y := 0; y < MatrixDimension; y++ {
			if !matchMatrix[y][x] {
				count += board[y][x]
			}
		}
	}
	return count * number
}

func solveBingo(fileName string) {
	bingoPicks, bingoBoard, matchMatrix := parseInputFile(fileName)

	scorePart1 := -1
	scorePart2 := -1
	boardsWon := 0
	boardCompleted := make([]bool, len(bingoBoard))

	for _, pick := range bingoPicks {
		if scorePart1 != -1 && scorePart2 != -1 {
			break
		}

		for i := range bingoBoard {
			match, y, x := isNumberonBoard(bingoBoard[i], pick)
			if match {
				matchMatrix[i][y][x] = true
			}
			if checkBingoBoardHorizontal(matchMatrix[i]) || checkBingoBoardVertical(matchMatrix[i]) {
				if !boardCompleted[i] {
					boardsWon++
					boardCompleted[i] = true
				}

				if scorePart1 == -1 {
					scorePart1 = computeScore(bingoBoard[i], matchMatrix[i], pick)
				}

				if boardsWon >= len(bingoBoard) && scorePart2 == -1 {
					scorePart2 = computeScore(bingoBoard[i], matchMatrix[i], pick)
				}
			}
		}
	}
	fmt.Println("Part1", scorePart1)
	fmt.Println("Part2", scorePart2)
}

func main() {

	// sanity check
	if MatrixDimension&1 != 1 {
		log.Fatalf("Matrix Dimension must be an odd number %v\n", MatrixDimension)
	}
	solveBingo("day-4-input.txt")
}
