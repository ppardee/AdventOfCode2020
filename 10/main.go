package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := []int64{}

	scanner := getScanner("input.txt")

	for scanner.Scan() {
		v, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		input = append(input, v)
	}
	j1Count, j3Count := getJumpCounts(input)
	fmt.Printf("The product of %v and %v is %v\n", j1Count, j3Count, j1Count*j3Count)

}

func getJumpCounts(input []int64) (int64, int64) {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	j1Count := int64(0)
	j3Count := int64(0)

	lastVal := int64(0)
	for _, v := range input {
		if v-lastVal == 1 {
			j1Count++
		} else if v-lastVal == 3 {
			j3Count++
		} else {
			panic("Invalid jolt jump!")
		}

		lastVal = v
	}

	return j1Count, j3Count + 1
}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}
