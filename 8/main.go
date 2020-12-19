package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := []command{}
	scanner := getScanner("input.txt")

	for scanner.Scan() {
		input = append(input, parseCommand(scanner.Text()))
	}

	acc := getAccumulatorFixInfiniteLoop(input)
	fmt.Printf("Final accumulator value: %v", acc)
}

func getAccumulatorAtInfiniteLoop(input []command) (int, bool) {
	hasRun := make([]bool, len(input))
	acc := 0
	curr := 0
	ct := 0
	for {
		ct++
		if hasRun[curr] {
			fmt.Printf("Found infinite loop after %v instructions at index: %v\n", ct, curr+1)
			return acc, true
		}

		hasRun[curr] = true
		cmd := input[curr]

		switch cmd.commandType {
		case "acc":
			acc += cmd.value
			curr++
		case "jmp":
			curr += cmd.value
		case "nop":
			curr++
			continue
		default:
			panic(fmt.Sprintf("Invalid command: %v\n", cmd.commandType))
		}

		if curr == len(input) {
			return acc, false
		}
	}
}

func getAccumulatorFixInfiniteLoop(input []command) int {

	// this could be made more efficient by walking the command tree and only touching the NEXT nop or jmp command.
	// As it is now, we may be altering commands that would not have been hit before the infinite loop bug
	for i := 0; i < len(input); i++ {

		if input[i].commandType == "acc" {
			continue
		}
		alter := make([]command, len(input))
		copy(alter, input)
		switch alter[i].commandType {
		case "jmp":
			alter[i].commandType = "nop"
			fmt.Printf("Changing the jmp at index %v to nop\n", i+1)
		case "nop":
			alter[i].commandType = "jmp"
			fmt.Printf("Changing the nop at index %v to jmp\n", i+1)
		default:
			panic("Invalid command type")
		}

		ret, infinite := getAccumulatorAtInfiniteLoop(alter)
		if !infinite {
			return ret
		}

	}

	return 0
}

func parseCommand(cmd string) command {
	var parts = strings.Split(cmd, " ")
	ret := command{
		commandType: parts[0],
	}

	v, _ := strconv.ParseInt(parts[1], 10, 64)

	ret.value = int(v)

	return ret

}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}

type command struct {
	commandType string
	value       int
}
