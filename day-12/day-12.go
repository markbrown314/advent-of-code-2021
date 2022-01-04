package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"

	"github.com/markbrown314/advent-of-code-2021/graph"
)

/*
    start
    /   \
c--A-----b--d
    \   /
     end
*/
func traverseCave(g graph.Graph, id string, visited []string) int {
	traverseCount := 0
	if visited == nil {
		visited = make([]string, 0)
	}

	//fmt.Printf("enter %v visited %v\n", id, visited)

	// mark as visted
	visited = append(visited, id)

	if id == "end" {
		fmt.Printf("** %v\n", visited)
		return 1
	}

	// copy slice of visited id's to map for faster lookup
	visitMap := make(map[string]bool)
	for _, v := range visited {
		if unicode.IsLower(rune(v[0])) {
			//fmt.Printf("marked %v\n", v)
			visitMap[v] = true
		}
	}

	for v := range g.Vertices[id].Edges {
		//fmt.Printf("candidate %v->%v\n", id, v)
		_, reached := visitMap[v]
		if !reached {
			//fmt.Printf("accepted %v\n", v)
			traverseCount += traverseCave(g, v, visited)
			visitMap[v] = true
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
		graph.AddEdgeToGraph(g, node[0], node[1], false)
	}
	fmt.Println(g.Vertices)
	traverseCount := traverseCave(g, "start", nil)
	fmt.Printf("part 1: traverseCount %v\n", traverseCount)
}
