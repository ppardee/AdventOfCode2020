package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//https://adventofcode.com/2020/day/4

	scanner := getScanner("input.txt")
	passports := make([]passport, 1)
	buffer := bytes.NewBufferString("")
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			buffer.WriteString(" " + line)
		} else {
			passports = addPassport(buffer, passports)
			buffer.Reset()
		}
	}

	passports = addPassport(buffer, passports[:])

	validCount := 0
	for _, v := range passports {
		if isValid(&v) {
			validCount++
		}
	}

	fmt.Println("Number of passports", len(passports))
	fmt.Println("Number of valid passports", validCount)

}

func isValid(passport *passport) bool {
	return yearIsValid(passport.birthYear, 1920, 2002) &&
		yearIsValid(passport.issueYear, 2010, 2020) &&
		yearIsValid(passport.expirationYear, 2020, 2030) &&
		hairColorIsValid(passport.hairColor) &&
		heightIsValid(passport.height) &&
		eyeColorIsValid(passport.eyeColor) &&
		pidIsValid(passport.pid)
}

func yearIsValid(v string, min int, max int) bool {
	if len(v) != 4 {
		return false
	}

	yr, err := strconv.ParseInt(v, 0, 0)
	if err != nil {
		return false
	}

	return yr >= int64(min) && yr <= int64(max)
}

func heightIsValid(v string) bool {
	length := len(v)
	if length < 4 {
		return false
	}
	if !(strings.HasSuffix(v, "cm") || strings.HasSuffix(v, "in")) {
		return false
	}

	var trunc string = v[:length-2]
	var suffix string = v[length-2:]
	num, err := strconv.ParseInt(trunc, 0, 0)
	if err != nil {
		return false
	}

	if suffix == "cm" {
		return num >= 150 && num <= 193
	}

	return num >= 59 && num <= 76

}

func hairColorIsValid(v string) bool {
	if len(v) != 7 {
		return false
	}

	pattern := `^#([a-f0-9]{6})$`
	matched, err := regexp.MatchString(pattern, v)

	if err != nil {
		fmt.Printf("Hair color error: %v", err)
		return false
	}

	return matched

}
func eyeColorIsValid(v string) bool {

	if len(v) != 3 {
		return false
	}
	validChoices := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, c := range validChoices {
		if c == v {
			return true
		}
	}

	return false
}

func pidIsValid(v string) bool {
	fmt.Printf("pid: %v ", v)
	if len(v) != 9 {
		fmt.Println("is invalid")
		return false
	}

	// These numbers have leading zeros, so they are interpreted as octal.  We need to specifically call them out as decimal
	_, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		fmt.Printf("is invalid because the parse failed - %v", err)
		return false
	}
	fmt.Println("is valid")
	return true
}

func addPassport(buf *bytes.Buffer, passports []passport) []passport {
	str := buf.String()
	pass := buildPassport(str)
	return append(passports, pass)
}

func buildPassport(input string) passport {

	res := passport{}

	a := strings.Split(input, " ")
	for _, v := range a {
		if len(v) == 0 {
			continue
		}

		kvp := strings.Split(v, ":")

		if len(kvp) != 2 {
			panic("Expecting key value pair")
		}

		switch kvp[0] {
		case "byr":
			res.birthYear = kvp[1]
		case "iyr":
			res.issueYear = kvp[1]
		case "eyr":
			res.expirationYear = kvp[1]
		case "hgt":
			res.height = kvp[1]
		case "hcl":
			res.hairColor = kvp[1]
		case "ecl":
			res.eyeColor = kvp[1]
		case "pid":
			res.pid = kvp[1]
		case "cid":
			res.cid = kvp[1]
		default:
			panic("Unexpected key")
		}

	}

	return res
}

func getScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	return bufio.NewScanner(reader)
}

type passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	pid            string
	cid            string
}
