package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//https://adventofcode.com/2020/day/4

	fileContents := getContents("input.txt")

	matrix := buildBagMatrix(fileContents)

	sum := 0
	for _, v := range matrix {
		if v {
			sum++
		}
	}

	fmt.Printf("The sum is %v\n", sum)
}

func buildBagMatrix(fileContents string) map[string]bool {
	retVal := make(map[string]bool)

	allBags := make(map[string][]string)

	direct := make(map[string]bool)

	lines := strings.Split(fileContents, "\n")
	for _, line := range lines {

		container, contains := parseLine(line)

		if _, ok := allBags[container]; ok {
			panic(fmt.Sprintf("Found a duplicate bag type %v", container))
		}

		allBags[container] = contains

		if canContainShinyGoldBag(contains) {
			direct[container] = true
		}
	}
	previousCount := 0
	for {
		previousCount = len(retVal)
	top:
		for k, v := range allBags {

			if direct[k] {
				retVal[k] = true
				continue
			}

			for _, b := range v {
				if direct[b] || retVal[b] {
					retVal[k] = true
					continue top
				}
			}
		}

		if previousCount == len(retVal) {
			break
		}

	}

	return retVal
}
func canContainShinyGoldBag(contains []string) bool {
	for _, v := range contains {
		if v == "shiny gold" {
			return true
		}
	}

	return false
}

func parseLine(line string) (string, []string) {
	container := ""
	contains := make([]string, 0)

	delim := " bags contain "

	i := strings.Index(line, delim)
	container = line[:i]

	containsStr := line[i+len(delim):]

	for _, v := range strings.Split(containsStr, ", ") {
		if "no other bags." == v {
			break
		}

		parts := strings.Split(v, " ")

		contains = append(contains, parts[1]+" "+parts[2])
	}

	return container, contains

}

func getContents(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	s := string(file)

	return s
}
