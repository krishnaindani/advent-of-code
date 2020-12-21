package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

var (
	birthYear      = "byr"
	issueYear      = "iyr"
	expirationyear = "eyr"
	height         = "hgt"
	hairColor      = "hcl"
	eyeColor       = "ecl"
	passportID     = "pid"
)

func main() {
	data := readFile()

	passports := parsePassport(data)

	validPassportsCount := validPassports(passports)

	fmt.Println("validPassports", validPassportsCount)
}

func validPassports(passports []string) int {
	fmt.Println("lenght", len(passports))
	var validPassports int
	for _, p := range passports {
		validCheck := validPassportFieldContains(p)
		if validCheck {
			valid := validPassport(p)
			if valid {
				validPassports++
			}
		}
	}
	return validPassports
}

func validPassportFieldContains(passport string) bool {
	for _, field := range mandatoryFields {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func validPassport(passport string) bool {
	passportFields := strings.Split(passport, " ")
	for _, f := range passportFields {
		fields := strings.Split(f, ":")
		if len(f) > 2 {
			field := fields[0]
			value := fields[1]
			switch field {
			case birthYear:
				y, _ := strconv.Atoi(value)
				valid := birthYearValidation(y)
				if !valid {
					return false
				}
			case issueYear:
				y, _ := strconv.Atoi(value)
				valid := issueYearValidation(y)
				if !valid {
					return false
				}
			case expirationyear:
				y, _ := strconv.Atoi(value)
				valid := expirationYearValidation(y)
				if !valid {
					return false
				}
			case height:
				valid := heightValidation(value)
				if !valid {
					return false
				}
			case hairColor:
				valid := hairColorValidation(value)
				if !valid {
					return false
				}
			case eyeColor:
				valid := eyeColorValiadation(value)
				if !valid {
					return false
				}
			case passportID:
				valid := passportIDValidation(value)
				if !valid {
					return false
				}
			}
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

	file, _ := os.Open("./data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	input = append(input, "")
	return input
}

func birthYearValidation(birthYear int) bool {
	return ((birthYear >= 1920) && (birthYear <= 2002))
}

func issueYearValidation(issueYear int) bool {
	return ((issueYear >= 2010) && (issueYear <= 2020))
}

func expirationYearValidation(expirationYear int) bool {
	return ((expirationYear >= 2020) && (expirationYear <= 2030))
}

func eyeColorValiadation(color string) bool {

	return (color == "amb") || (color == "blu") || (color == "gry") ||
		(color == "brn") || (color == "hzl") || (color == "oth") ||
		(color == "grn")
}

func passportIDValidation(id string) bool {
	return len(id) == 9
}

func hairColorValidation(color string) bool {
	match, _ := regexp.MatchString("^#([0-9a-f]{6})", color)
	return match
}

func heightValidation(height string) bool {

	if strings.Contains(height, "cm") {
		value, _ := strconv.Atoi(strings.Split(height, "cm")[0])
		return (value >= 150) && (value <= 193)
	}

	if strings.Contains(height, "in") {
		value, _ := strconv.Atoi(strings.Split(height, "in")[0])
		return (value >= 59) && (value <= 76)
	}

	return false
}
