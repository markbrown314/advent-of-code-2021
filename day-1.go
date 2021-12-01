package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	input, err := ioutil.ReadFile("day-1-input-1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	sliceInput := strings.Split(string(input), "\n")
	fmt.Println(sliceInput)

	prev := 0
	increased := 0

	for i := 0; i < len(sliceInput); i++ {
		depth, err := strconv.Atoi(sliceInput[i])
		if err != nil {
			log.Fatalf("unable to convert input: %v", err)
		}
		if prev == 0 {
			fmt.Printf("%v (N/A - no previous measurement)\n", depth)
			prev = depth
			continue
		}

		fmt.Print(depth, " ")

		if depth == prev {
			fmt.Print("(no change)")
		}

		if depth > prev {
			fmt.Print("(increased)")
			increased++
			prev = depth
		}

		if depth < prev {
			fmt.Print("(decreased)")
			prev = depth
		}

		fmt.Println()
	}

	fmt.Printf("depth has increased %v times", increased)
}
