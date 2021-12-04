package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func pilotSubmarine(commandList []string, enableAim bool) int {
	command_mode := true
	command := "none"
	position := 0
	depth := 0
	aim := 0

	for _, input := range commandList {
		if command_mode {
			command = input
			fmt.Println("command:", command)
		} else {
			magnitude, err := strconv.Atoi(input)
			if err != nil {
				log.Fatalf("unable to convert input to int %v", input)
			}
			fmt.Println("magnitude:", magnitude)
			switch command {
			case "forward":
				position += magnitude
				if enableAim {
					depth += aim * magnitude
				}
			case "down":
				if !enableAim {
					depth += magnitude
				}
				aim += magnitude
			case "up":
				if !enableAim {
					depth -= magnitude
				}
				aim -= magnitude
			default:
				log.Fatalf("unknown command %v", command)
			}
			fmt.Println("aim:", aim, "position:", position, "depth:", depth)
		}
		command_mode = !command_mode
	}
	fmt.Printf("position %v and depth %v\n", position, depth)
	return position * depth
}

func main() {
	input, err := ioutil.ReadFile("day-2-input-1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	fmt.Println("part 1:", pilotSubmarine(strings.Fields(string(input)), false))
	fmt.Println("part 2:", pilotSubmarine(strings.Fields(string(input)), true))
}
