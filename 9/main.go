package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	preambleSize := 25
	input := []int64{}

	scanner := getScanner("input.txt")

	for scanner.Scan() {
		v, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		input = append(input, v)
	}

	violator := int64(0)

	for i, v := range input {
		if i < preambleSize {
			continue
		}
		if !hasAddends(input[i-preambleSize:i], v) {
			violator = v
		}
	}

	fmt.Printf("violator: %v\n", violator)
	min, max := getViolatorSetMinMax(input, violator)

	fmt.Printf("The sum of %v and %v is %v\n", min, max, min+max)

}

func getViolatorSetMinMax(input []int64, violator int64) (int64, int64) {
	min := int64(0)
	max := int64(1)

	for {
		sum := int64(0)
		for i := min; i <= max; i++ {
			sum += input[i]
		}
		if sum < violator {
			max++
		} else if sum > violator {
			min++
		} else if sum == violator {
			return findFloorAndCeiling(input[min:max])

		}
	}
}

func findFloorAndCeiling(input []int64) (int64, int64) {
	floor := input[0]
	ceiling := input[0]

	for _, v := range input {
		if v < floor {
			floor = v
		}

		if v > ceiling {
			ceiling = v
		}

	}

	return floor, ceiling
}

func hasAddends(segment []int64, target int64) bool {
	for i := 0; i < len(segment)-1; i++ {
		for j := i + 1; j < len(segment); j++ {
			if segment[i]+segment[j] == target {
				return true
			}
		}
	}
	return false
}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}
