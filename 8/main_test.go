package main

import "testing"

func TestGetAccumulatorAtInfiniteLoop(t *testing.T) {
	testData := []command{
		parseCommand("nop +0"),
		parseCommand("acc +1"),
		parseCommand("jmp +4"),
		parseCommand("acc +3"),
		parseCommand("jmp -3"),
		parseCommand("acc -99"),
		parseCommand("acc +1"),
		parseCommand("jmp -4"),
		parseCommand("acc +6"),
	}

	val, _ := getAccumulatorAtInfiniteLoop(testData)
	if val != 5 {
		t.Errorf("Expected 5 but got %v", val)
	}
}

func TestGetAccumulatorFixInfiniteLoop(t *testing.T) {
	testData := []command{
		parseCommand("nop +0"),
		parseCommand("acc +1"),
		parseCommand("jmp +4"),
		parseCommand("acc +3"),
		parseCommand("jmp -3"),
		parseCommand("acc -99"),
		parseCommand("acc +1"),
		parseCommand("jmp -4"),
		parseCommand("acc +6"),
	}

	val := getAccumulatorFixInfiniteLoop(testData)
	if val != 8 {
		t.Errorf("Expected 8 but got %v", val)
	}
}
