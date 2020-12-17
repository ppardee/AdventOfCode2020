package main

import "testing"

func TestGetSeatID(t *testing.T) {

	testData := map[string]int64{
		"BFFFBBFRRR": 567,
		"FFFBBBFRRR": 119,
		"BBFFBBFRLL": 820,
		"FBFBBFFRLR": 357,
	}

	for k, v := range testData {
		id := getSeatID(k)
		if id != v {
			t.Errorf("Failed for %v: Expected %v but got %v", k, v, id)
		}
	}

}

func TestFindMissingSeat(t *testing.T) {

	testData := []int{
		9,
		10,
		5,
		6,
		8,
		11,
	}

	id := findMissingSeat(testData)
	if id != 7 {
		t.Errorf("Expected 7, got %v", id)
	}

}
