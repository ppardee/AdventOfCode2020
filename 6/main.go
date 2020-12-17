package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//https://adventofcode.com/2020/day/4

	fileContents := getContents("input.txt")

	groupStrings := getGroupStrings(fileContents)

	sum := 0

	for _, v := range groupStrings {
		sum = sum + getGroupUniqueCount(v)
	}

	fmt.Printf("The sum of unique values is %v\n", sum)

	sum = 0

	for _, v := range groupStrings {
		sum = sum + getGroupUnanimousCount(v)
	}

	fmt.Printf("The sum of unanimous values is %v\n", sum)
}

func getGroupStrings(allInput string) []string {
	res := make([]string, 0)
	b := bytes.NewBufferString("")

	for _, v := range strings.Split(allInput, "\n") {
		trimmedV := strings.TrimSpace(v)
		if len(trimmedV) > 0 {
			b.WriteString(trimmedV)
			b.WriteString("\n")
			continue
		}
		res = append(res, b.String())
		b.Reset()
	}

	// Catch the tail
	res = append(res, b.String())

	return res
}

func getGroupUniqueCount(group string) int {
	// Counts how many 'yes' answers there are in the group
	// there are 26 answers, a-z.  Duplicates don't count
	answers := make(map[rune]bool)

	for _, v := range group {
		if v == '\n' {
			continue
		}
		answers[v] = true
	}

	return len(answers)
}

func getGroupUnanimousCount(group string) int {
	// Counts how many answers everyone said yes to
	answers := make(map[rune]int)

	memberCount := 0

	for _, v := range group {
		if v == '\n' {
			memberCount++
			continue
		}
		answers[v]++
	}

	unanimousCount := 0
	for _, v := range answers {
		if v == memberCount {
			unanimousCount++
		}
	}
	return unanimousCount
}

func getContents(fileName string) string {
	file, _ := ioutil.ReadFile(fileName)
	return string(file)
}
