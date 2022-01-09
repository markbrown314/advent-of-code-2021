package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

type PolymerRuleMap struct {
	RuleMap map[string]string
}

func initPolymerRuleMap() PolymerRuleMap {
	template := PolymerRuleMap{RuleMap: make(map[string]string)}
	return template
}

func (r *PolymerRuleMap) runPolymerRules(template string, iterations int) string {
	var polymer string
	var rule string
	for step := 0; step < iterations; step++ {
		polymer = ""
		for i := 0; i < len(template)-1; i++ {
			rule = template[i : i+2]
			insertion, valid := r.RuleMap[rule]
			if !valid {
				log.Fatalf("error could not find rule %v", rule)
			}
			polymer += (string(rule[0]) + insertion)
		}
		if len(rule) != 2 {
			log.Fatalf("error running rule map\n")
		}
		// add last letter
		polymer += string(rule[1])
		template = polymer
	}
	return polymer
}

func solvePart1(polymer string) int {
	letterMap := make(map[rune]int)
	freq := make([]int, 0)
	for _, l := range polymer {
		letterMap[l]++
	}
	for l := range letterMap {
		freq = append(freq, letterMap[l])
	}

	sort.Ints(freq)
	return freq[len(freq)-1] - freq[0]
}

func main() {
	fmt.Println("Problem Day #14")
	fileInput, err := ioutil.ReadFile("day-14-input.txt")
	if err != nil {
		log.Fatalf("error loading file %v\n", err)
	}

	insertionRuleRe := regexp.MustCompile("[A-Z][A-Z] -> [A-Z]")

	lines := strings.Split(string(fileInput), "\n")
	template := lines[0]
	r := initPolymerRuleMap()

	for _, rule := range lines {
		if insertionRuleRe.MatchString(rule) {
			insertionRule := strings.Split(rule, " -> ")
			r.RuleMap[insertionRule[0]] = insertionRule[1]
		}
	}
	polymer := r.runPolymerRules(template, 10)
	fmt.Printf("part 1: result %v\n", solvePart1(polymer))
}
