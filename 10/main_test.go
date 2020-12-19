package main

import "testing"

func TestGetJumpCounts(t *testing.T) {
	testData := []int64{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}

	j1, j3 := getJumpCounts(testData)

	if j1 != 22 {
		t.Errorf("Expected j1 == 22 but got %v", j1)
	}

	if j3 != 10 {
		t.Errorf("Expected j3 == 10 but got %v", j3)
	}

}
