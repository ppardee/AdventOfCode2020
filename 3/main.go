package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//https://adventofcode.com/2020/day/3

	scanner := getScanner("input.txt")
	var treeMap [323][31]bool
	line := 0
	for scanner.Scan() {
		for i, v := range scanner.Text() {
			treeMap[line][i] = v == '#'
		}
		line++
	}

	treeProduct := treesOnSlope(1, 1, &treeMap)
	treeProduct = treeProduct * treesOnSlope(3, 1, &treeMap) // Part one
	treeProduct = treeProduct * treesOnSlope(5, 1, &treeMap)
	treeProduct = treeProduct * treesOnSlope(7, 1, &treeMap)
	treeProduct = treeProduct * treesOnSlope(1, 2, &treeMap)
	fmt.Println("Tree Product = ", treeProduct)

}

func treesOnSlope(x int, y int, treeMap *[323][31]bool) int {
	treeCount := 0

	for i, j := 0, 0; i < 323; {
		if treeMap[i][j] {
			treeCount++
		}

		i = i + y
		j = j + x

		if j >= len(treeMap[i]) {
			j = j - len(treeMap[i])
		}
	}
	return treeCount
}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}
