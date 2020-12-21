package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	minByr          = 1920
	maxByr          = 2002
	minIyr          = 2010
	maxIyr          = 2020
	minEyr          = 2020
	maxEyr          = 2030
)

func main() {
	data := readFile()
	fmt.Println("data", data)

	passports := parsePassport(data)
	validPassports := 0
	for _, p := range passports {
		valid := checkValidPassports(p)
		if valid {
			validPassports++
		}
	}
	fmt.Println("validPassports", validPassports)
}

func checkValidPassports(passport string) bool {
	for _, field := range mandatoryFields {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func parsePassport(passports []string) []string {
	parsedPassports := []string{}
	passport := ""
	for _, p := range passports {
		if len(p) > 0 {
			passport += p
		} else {
			parsedPassports = append(parsedPassports, passport)
			passport = ""
		}
	}
	return parsedPassports
}

func readFile() []string {
	var input []string

	file, _ := os.Open("./sampleData.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	input = append(input, "")
	return input
}

func validateBirthYear(y int) bool {
	return ((y >= minByr) && (y <= maxByr))
}

func validateIssueYear(y int) bool {
	return ((y >= minIyr) && (y <= maxIyr))
}
