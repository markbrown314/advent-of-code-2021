package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/markbrown314/advent-of-code-2021/graph"
)

/*
    start
    /   \
c--A-----b--d
    \   /
     end
*/

func main() {
	fmt.Println("Problem Day #12")
	fileInput, err := ioutil.ReadFile("day-12-test1.txt")
	if err != nil {
		log.Fatalf("error: failed to read file %v\n", err)
	}

	g := graph.New()

	for _, line := range strings.Split(string(fileInput), "\n") {
		if line == "" {
			continue
		}
		node := strings.Split(line, "-")
		graph.AddEdgeToGraph(g, node[0], node[1], true)
	}
	fmt.Println(g.Vertices)
}
