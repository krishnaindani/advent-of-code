package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	data := readFile()

	passports := parsePassport(data)

	validPassportsCount := validPassports(passports)

	fmt.Println("validPassports", validPassportsCount)
}

func validPassports(passports []string) int {
	var validPassports int
	for _, p := range passports {
		if validPassport(p) {
			validPassports++
		}
	}
	return validPassports
}

func validPassport(passport string) bool {
	for _, field := range mandatoryFields {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func parsePassport(passports []string) []string {
	var parsedPassports []string
	var passport string
	for _, p := range passports {
		if len(p) > 0 {
			passport += p
			passport += " "
		} else {
			parsedPassports = append(parsedPassports, passport)
			passport = ""
		}
	}
	return parsedPassports
}

func readFile() []string {
	var input []string

	file, err := os.Open("./sampledata.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	input = append(input, "")
	return input
}
