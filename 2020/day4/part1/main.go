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
		valid := validPassport(p)
		if valid {
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
