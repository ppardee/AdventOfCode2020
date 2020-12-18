package main

import (
	"testing"
)

func TestBuildBagMatrix(t *testing.T) {
	testData := "light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellowbags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dottedblack bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags.\nfat ass bags contain 1 light red bag."

	expected := map[string]bool{
		"light red":    true,
		"dark orange":  true,
		"bright white": true,
		"muted yellow": true,
		"shiny gold":   false,
		"dark olive":   false,
		"vibrant plum": false,
		"faded blue":   false,
		"dotted black": false,
		"fat ass":      true,
	}

	res := buildBagMatrix(testData)
	if len(res) == 0 {
		t.Errorf("buildBagMatrix returned no results")
	}
	for k, v := range res {
		e := expected[k]
		if e != v {
			t.Errorf("For %v - Expected %v, but got %v", k, e, v)
		}
	}

}

// func TestCanContainGoldBag(t *testing.T) {
// 	testData := map[string]bool{
// 		"light red bags contain 1 bright white bag, 2 muted yellow bags.":    true,
// 		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.": true,
// 		"bright white bags contain 1 shiny gold bag.":                        true,
// 		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.":    true,
// 		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.":     false,
// 		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.":    false,
// 		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.":  false,
// 		"faded blue bags contain no other bags.":                             false,
// 		"dotted black bags contain no other bags.":                           false,
// 	}
// 	bagMatrix := map[string]bool{
// 		"light red":    true,
// 		"dark orange":  true,
// 		"bright white": true,
// 		"muted yellow": true,
// 		"shiny gold":   false,
// 		"dark olive":   false,
// 		"vibrant plum": false,
// 		"faded blue":   false,
// 		"dotted black": false,
// 	}
// 	for k, v := range testData {
// 		canContain := canContainGoldBag(bagMatrix, k)
// 		if canContain != v {
// 			t.Errorf("For %v - Expected %v, but got %v", k, v, canContain)
// 		}
// 	}
// }

func TestParseLine(t *testing.T) {
	testData := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	expected := map[string][]string{
		"light red":    []string{"bright white", "muted yellow"},
		"dark orange":  []string{"bright white", "muted yellow"},
		"bright white": []string{"shiny gold"},
		"muted yellow": []string{"shiny gold", "faded blue"},
		"shiny gold":   []string{"dark olive", "vibrant plum"},
		"dark olive":   []string{"faded blue", "dotted black"},
		"vibrant plum": []string{"faded blue", "dotted black"},
		"faded blue":   []string{},
		"dotted black": []string{},
	}

	for _, v := range testData {
		key, val := parseLine(v)

		if len(key) == 0 {
			t.Errorf("For %v - Empty key (container) returned", v)
		}

		e := expected[key]

		if len(e) != len(val) {
			t.Errorf("For %v - Expected %v bags and got %v - %v", v, len(e), len(val), val)
		}

		for idx, c := range val {
			if e[idx] != c {
				t.Errorf("For %v index %v - Expected %v and got %v", v, idx, e[idx], c)
			}
		}

	}
}
