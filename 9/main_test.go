package main

import "testing"

func TestHasAddends(t *testing.T) {
	testData := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	violator := int64(0)
	for i, v := range testData {
		if i < 5 {
			continue
		}
		if !hasAddends(testData[i-5:i], v) {
			violator = v
		}
	}

	if violator != 127 {
		t.Errorf("Expected 127 but got %v", violator)
	}

}

func TestGetViolatorSetMinMax(t *testing.T) {
	testData := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	violator := int64(127)
	min, max := getViolatorSetMinMax(testData, violator)

	if min != 15 {
		t.Errorf("Expected min 15 but got %v", min)
	}

	if max != 47 {
		t.Errorf("Expected max 47 but got %v", max)
	}
}
