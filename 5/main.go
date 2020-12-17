package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	//https://adventofcode.com/2020/day/4

	scanner := getScanner("input.txt")

	highestID := int64(0)
	seats := make([]int, 0)

	for scanner.Scan() {
		id := getSeatID(scanner.Text())
		if id > highestID {
			highestID = id
		}

		seats = append(seats, int(id))

	}

	fmt.Printf("Highest ID is %v\n", highestID)
	fmt.Printf("Your seat number is %v\n", findMissingSeat(seats))
}

func findMissingSeat(seats []int) int {

	sort.Ints(seats)

	lastVal := seats[0] - int(1)

	for _, v := range seats {
		if lastVal+1 == v {
			lastVal = v
			continue
		}
		break
	}
	return lastVal + 1
}

func getSeatID(code string) int64 {
	// This is a binary tree search
	// For rows
	// F means take lower half
	// B means take upper half

	// For columns
	// R means take upper half
	// L means take lower half

	colUpper, colLower := 8, 1
	rowUpper, rowLower := 128, 1

	//fmt.Printf("Decoding %v \n", code)

	for _, v := range code {
		switch v {
		case 'F':
			shift := (rowUpper - rowLower + 1) / 2
			rowUpper = rowUpper - shift
			//fmt.Printf("Found 'F'.  Reducing rowUpper to %v \n", rowUpper)

		case 'B':
			shift := (rowUpper - rowLower + 1) / 2
			rowLower = rowLower + shift
			//fmt.Printf("Found 'B'.  Increasing rowLower to %v \n", rowLower)

		case 'L':
			shift := (colUpper - colLower + 1) / 2
			colUpper = colUpper - shift
			//fmt.Printf("Found 'L'.  Reducing colUpper to %v \n", colUpper)
		case 'R':
			shift := (colUpper - colLower + 1) / 2
			colLower = colLower + shift
			//fmt.Printf("Found 'R'.  Increasing colLower to %v \n", colLower)
		}
	}

	// At this point the upper and lower bounds should be the same, but we need to shift them down by one because the puzzle is 0-based
	colUpper = colUpper - 1
	rowUpper = rowUpper - 1

	seatNumber := int64(rowUpper*8) + int64(colUpper)

	//fmt.Printf("Seat number is %v \n", seatNumber)

	return seatNumber

}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}
