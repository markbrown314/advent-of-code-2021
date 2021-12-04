package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	input, err := ioutil.ReadFile("day-3-test.txt")
	if err != nil {
		log.Fatalf("error reading input file %v", err)
	}

	lines := strings.Split(string(input), "\n")
	fmt.Println(len(lines))

}
