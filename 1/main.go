package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const target int = 2020

func findTwo() {

	scanner := getScanner("input.txt")

	numbers := getSortedInput(scanner)

	fmt.Println(numbers)

	highIdx := len(numbers) - 1
	lowIdx := 0

top:
	for ; lowIdx < highIdx; lowIdx++ {
		fmt.Printf("Checking %v ", numbers[lowIdx])

		for highIdx > lowIdx {
			fmt.Printf("against %v - ", numbers[highIdx])
			t := numbers[lowIdx] + numbers[highIdx]
			fmt.Printf("The sum is %v\n", t)
			if t == target {
				fmt.Println("Found it!")

				break top
			}

			if t > target {
				fmt.Println("Too high, next high number down!")
				highIdx--
				continue
			}

			if t < target {
				fmt.Println("Too low, next low number up!")
				break
			}
		}
	}

	fmt.Printf("The two numbers are %v and %v\n", numbers[lowIdx], numbers[highIdx])
	fmt.Printf("The product of these numbers is %v\n", numbers[lowIdx]*numbers[highIdx])
}

func findThree() {
	scanner := getScanner("input.txt")

	numbers := getSortedInput(scanner)

	fmt.Println(numbers)

	highIdx := 2
	midIdx := 1
	lowIdx := 0
top:
	for ; lowIdx < len(numbers)-3; lowIdx++ {
		for ; midIdx < highIdx; midIdx++ {
			for highIdx = midIdx + 1; highIdx < len(numbers); highIdx++ {
				l := numbers[lowIdx]
				m := numbers[midIdx]
				h := numbers[highIdx]
				sum := l + m + h

				if l+m+h == 2020 {
					fmt.Println("Found it!")
					break top
				}
				if sum < 2020 {
					fmt.Println("Too low, continuing!")
					continue
				}
				if sum > 2020 {
					fmt.Println("Too high, breaking!")
					break
				}

			}
		}
	}

	fmt.Printf("The three numbers are %v, %v and %v\n", numbers[lowIdx], numbers[midIdx], numbers[highIdx])
	fmt.Printf("The product of these numbers is %v\n", numbers[lowIdx]*numbers[midIdx]*numbers[highIdx])
}

func main() {
	findThree()
}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}

func getSortedInput(scanner *bufio.Scanner) []int {
	var numbers []int

	for scanner.Scan() {
		t := scanner.Text()
		n, err := strconv.ParseInt(t, 0, 0)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, int(n))
	}

	sort.Ints(numbers)
	return numbers
}
