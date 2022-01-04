package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"

	"github.com/markbrown314/advent-of-code-2021/graph"
)

func traverseCave(g graph.Graph, id string, visited []string, maxVisit int) int {
	traverseCount := 0
	if visited == nil {
		visited = make([]string, 0)
	} else {
		if id == "start" {
			return 0
		}
	}

	// mark as visted
	visited = append(visited, id)

	if id == "end" {
		return 1
	}

	// copy slice of visited id's to map for faster lookup
	visitMap := make(map[string]int)
	for _, v := range visited {
		if unicode.IsLower(rune(v[0])) {
			visitMap[v] += 1
		} else {
			visitMap[v] = 0
		}
		if visitMap[v] >= 2 {
			maxVisit = 1
		}
	}

	for v := range g.Vertices[id].Edges {
		visitCount := visitMap[v]
		if visitCount < maxVisit {
			traverseCount += traverseCave(g, v, visited, maxVisit)
			visitMap[v] += 1
		}
	}
	return traverseCount
}

func main() {
	fmt.Println("Problem Day #12")
	fileInput, err := ioutil.ReadFile("day-12-input.txt")
	if err != nil {
		log.Fatalf("error: failed to read file %v\n", err)
	}

	g := graph.New()

	for _, line := range strings.Split(string(fileInput), "\n") {
		if line == "" {
			continue
		}
		node := strings.Split(line, "-")
		g.AddEdge(node[0], node[1], false)
	}
	fmt.Printf("part 1: traverseCount %v\n", traverseCave(g, "start", nil, 1))
	fmt.Printf("part 2: traverseCount %v\n", traverseCave(g, "start", nil, 2))
}
