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

	//fmt.Printf("enter %v visited %v\n", id, visited)

	// mark as visted
	visited = append(visited, id)

	if id == "end" {
		//fmt.Printf("** %v\n", visited)
		for _, v := range visited {
			if v == "start" {
				fmt.Printf("start,")
				continue
			}
			if v == "end" {
				fmt.Printf("end\n")
				continue
			}
			fmt.Printf("%v,", v)
		}
		return 1
	}

	// copy slice of visited id's to map for faster lookup
	visitMap := make(map[string]int)
	for _, v := range visited {
		if unicode.IsLower(rune(v[0])) {
			//fmt.Printf("marked %v\n", v)
			visitMap[v] += 1
		} else {
			visitMap[v] = 0
		}
		if visitMap[v] >= 2 {
			maxVisit = 1
		}
	}

	for v := range g.Vertices[id].Edges {
		//fmt.Printf("candidate %v->%v\n", id, v)
		visitCount := visitMap[v]
		if visitCount < maxVisit {
			//fmt.Printf("accepted %v\n", v)
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
		graph.AddEdgeToGraph(g, node[0], node[1], false)
	}
	fmt.Println(g.Vertices)
	//fmt.Printf("part 1: traverseCount %v\n", traverseCave(g, "start", nil, 1))
	fmt.Printf("part 2: traverseCount %v\n", traverseCave(g, "start", nil, 2))
}
