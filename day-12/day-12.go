package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/markbrown314/advent-of-code-2021/utils"
)

func main() {
	fmt.Println("Problem Day #12")
	fileInput, err := ioutil.ReadFile("day-12-test1.txt")
	if err != nil {
		log.Fatalf("error: failed to read file %v\n", err)
	}

	g := utils.CreateDirGraph()

	for _, line := range strings.Split(string(fileInput), "\n") {
		if line == "" {
			continue
		}
		node := strings.Split(line, "-")
		utils.AddEdgeToGraph(g, node[0], node[1])
	}
	fmt.Println(g.Vertices)
}
