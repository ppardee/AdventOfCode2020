package main

import "testing"

func TestGetJumpCounts(t *testing.T) {
	testData := [][]string{
		{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", "L", "L", "L", "L", "L", "L", ".", "L", "L"},
		{"L", ".", "L", ".", "L", ".", ".", "L", ".", "."},
		{"L", "L", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
		{".", ".", "L", ".", "L", ".", ".", ".", ".", "."},
		{"L", "L", "L", "L", "L", "L", "L", "L", "L", "L"},
		{"L", ".", "L", "L", "L", "L", "L", "L", ".", "L"},
		{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
	}

	count := runSim(testData)
	if count != 37 {
		t.Errorf("Expected count == 37 but got %v", count)
	}

}
