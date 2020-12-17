package main

import (
	"testing"
)

func TestGetGroupStrings(t *testing.T) {

	testData := "sukewbthnjdyv\nsjubpeyvx\nouehaytkgzbv\nqfueimylvcrb\n\nfoyedvcm\nyfodemvc\ncydvomef\nfdvmocey\nyvefocmd\n\nyizvxsnrhafmulctdp\nzvpnhusixraejywfcmt\nfxaiustovzhnkmrypc\nbtcfymhrnapuvxwdisz"
	expected := []string{
		"sukewbthnjdyv\nsjubpeyvx\nouehaytkgzbv\nqfueimylvcrb\n",
		"foyedvcm\nyfodemvc\ncydvomef\nfdvmocey\nyvefocmd\n",
		"yizvxsnrhafmulctdp\nzvpnhusixraejywfcmt\nfxaiustovzhnkmrypc\nbtcfymhrnapuvxwdisz\n",
	}

	groupStrings := getGroupStrings(testData)

	if len(groupStrings) != 3 {
		t.Errorf("Expected 3 groups, but got %v", len(groupStrings))
	}

	for i, v := range groupStrings {
		if expected[i] != v {
			t.Errorf("Expected %v, but got %v", expected[i], v)
		}
	}

}

func TestGetGroupUniqueCount(t *testing.T) {
	testData := map[string]int{
		"sukewbthnjdyvsjubpeyvxouehaytkgzbvqfueimylvcrb":                             26,
		"foyedvcmyfodemvccydvomeffdvmoceyyvefocmd":                                   8,
		"yizvxsnrhafmulctdpzvpnhusixraejywfcmtfxaiustovzhnkmrypcbtcfymhrnapuvxwdisz": 24,
	}

	for k, v := range testData {
		count := getGroupUniqueCount(k)
		if count != v {
			t.Errorf("Expected %v, but got %v", v, count)
		}
	}
}

func TestGetGroupUniqueCount_WithNewLines(t *testing.T) {
	testData := map[string]int{
		"abc":     3,
		"a\nb\nc": 3,
		"ab\nac":  3,
	}

	for k, v := range testData {
		count := getGroupUniqueCount(k)
		if count != v {
			t.Errorf("Expected %v, but got %v", v, count)
		}
	}
}

func TestGetGroupUnanimousCount(t *testing.T) {
	// testData := map[string]int{
	// 	"sukewbthnjdyv\nsjubpeyvx\nouehaytkgzbv\nqfueimylvcrb\n":                             12,
	// 	"foyedvcm\nyfodemvc\ncydvomef\nfdvmocey\nyvefocmd\n":                                 15,
	// 	"yizvxsnrhafmulctdp\nzvpnhusixraejywfcmt\nfxaiustovzhnkmrypc\nbtcfymhrnapuvxwdisz\n": 18,
	// }
	testData := map[string]int{
		"abc\n":        3,
		"a\nb\nc\n":    0,
		"ab\nac\n":     1,
		"a\na\na\na\n": 1,
		"b\n":          1,
	}

	for k, v := range testData {
		count := getGroupUnanimousCount(k)
		if count != v {
			t.Errorf("Expected %v, but got %v", v, count)
		}
	}
}

func TestSummingItUp(t *testing.T) {
	testData := "sukewbthnjdyv\nsjubpeyvx\nouehaytkgzbv\nqfueimylvcrb\n\nfoyedvcm\nyfodemvc\ncydvomef\nfdvmocey\nyvefocmd\n\nyizvxsnrhafmulctdp\nzvpnhusixraejywfcmt\nfxaiustovzhnkmrypc\nbtcfymhrnapuvxwdisz"
	groupStrings := getGroupStrings(testData)

	sum := 0

	for _, v := range groupStrings {
		sum = sum + getGroupUniqueCount(v)
	}

	if sum != 58 {
		t.Errorf("Expected 58, but got %v", sum)
	}
}
