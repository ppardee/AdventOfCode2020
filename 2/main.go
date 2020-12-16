package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//https://adventofcode.com/2020/day/2

	validCount := 0

	scanner := getScanner("input.txt")

	for scanner.Scan() {

		a := strings.FieldsFunc(scanner.Text(), split)
		pos1, _ := strconv.ParseInt(a[0], 0, 0)
		pos2, _ := strconv.ParseInt(a[1], 0, 0)
		char := a[2][:1]
		password := a[3]

		// Doing -1 on the positions because they are not zero-based in the puzzle data
		if isValid2(pos1-1, pos2-1, char, password) {
			validCount++
		}
	}

	fmt.Println("Valid count = ", validCount)

}

func partOne() {
	validCount := 0

	scanner := getScanner("input.txt")

	for scanner.Scan() {

		a := strings.FieldsFunc(scanner.Text(), split)
		min, _ := strconv.ParseInt(a[0], 0, 0)
		limit, _ := strconv.ParseInt(a[1], 0, 0)
		char := a[2][:1]
		password := a[3]

		if isValid1(min, limit, char, password) {
			validCount++
		}
		//fmt.Println("", min, limit, char, password, valid)
	}

	fmt.Println("Valid count = ", validCount)
}

func isValid1(min int64, limit int64, char string, password string) bool {
	//2-6 c: fcpwjqhcgtffzlbj
	var count int64 = 0
	for _, c := range password {
		if c == rune(char[0]) {
			count++
		}
	}

	return count >= min && count <= limit
}

func isValid2(pos1 int64, pos2 int64, char string, password string) bool {
	//2-6 c: fcpwjqhcgtffzlbj

	pos1Contains := password[pos1] == char[0]
	pos2Contains := password[pos2] == char[0]
	isValid := (pos1Contains || pos2Contains) && password[pos1] != password[pos2]
	if (pos1Contains || pos2Contains) && password[pos1] != password[pos2] {
		fmt.Printf("Valid: %v %v %v %v\n", pos1, pos2, char, password)
	} else {
		fmt.Printf("Invalid: %v %v %v %v\n", pos1, pos2, char, password)
	}

	return isValid

}

func split(r rune) bool {
	return r == '-' || r == ' '
}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}
