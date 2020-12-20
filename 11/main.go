package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := make([][]string, 0)

	scanner := getScanner("input.txt")

	for scanner.Scan() {
		v := scanner.Text()
		col := make([]string, len(v))
		for i, v := range v {
			col[i] = string(v)
		}

		grid = append(grid, col)

	}
	count := runSim(grid)

	fmt.Printf("Number of occupied seats is %v \n", count)

}

func runSim(grid [][]string) int {

	prev := copyGrid(grid)

	for {
		// printGrid(prev)
		// fmt.Println("--------------------------------")
		new := step(prev)
		if slicesAreEqual(prev, new) {
			break
		}
		prev = copyGrid(new)
	}

	count := 0
	for i := range prev {
		for _, v := range prev[i] {
			if v == "#" {
				count++
			}
		}
	}

	return count
}

func printGrid(grid [][]string) {
	for i := range grid {
		for _, v := range grid[i] {
			fmt.Printf("%v", v)
		}
		fmt.Println()
	}
}

func slicesAreEqual(a [][]string, b [][]string) bool {
	for i := range a {
		for j, v := range b[i] {

			if a[i][j] != v {
				return false
			}

		}
	}

	return true
}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}

func step(grid [][]string) [][]string {
	ret := copyGrid(grid)

	for rowIdx := range grid {
		for colIdx, val := range grid[rowIdx] {
			switch val {
			case ".":
				continue
			case "L":
				// Empty seat: If there are no occupied seats adjacent to it, the seat becomes occupied.
				if getAdjacentSeatOccupiedCount(grid, rowIdx, colIdx) > 0 {
					continue
				}
				ret[rowIdx][colIdx] = "#"
			case "#":
				// Occupied Seat: If 4 or more adjacent seats are occupied, the seat becomes empty
				if getAdjacentSeatOccupiedCount(grid, rowIdx, colIdx) < 4 {
					continue
				}
				ret[rowIdx][colIdx] = "L"
			default:
				panic("Unexpected value")
			}

		}

	}
	return ret

}

func getAdjacentSeatOccupiedCount(grid [][]string, rowIdx int, colIdx int) int {
	ret := 0
	for i := rowIdx - 1; i <= rowIdx+1; i++ {
		if i < 0 || i >= len(grid) { // out of bounds
			continue
		}
		for j := colIdx - 1; j <= colIdx+1; j++ {

			if j < 0 || j >= len(grid[i]) { // out of bounds
				continue
			}

			if i == rowIdx && j == colIdx { // This is the cell we're testing
				continue
			}

			if grid[i][j] == "#" {
				ret++
			}
		}
	}

	return ret
}

func copyGrid(grid [][]string) [][]string {
	retVal := make([][]string, len(grid))

	for i, v := range grid {
		retVal[i] = make([]string, len(v))
		copy(retVal[i], v)
	}

	return retVal

}
