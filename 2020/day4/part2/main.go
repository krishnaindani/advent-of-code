package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/krishnaindani/advent-of-code/2020/day4/part2/validate"
)

var (
	birthYear      = "byr"
	issueYear      = "iyr"
	expirationYear = "eyr"
	height         = "hgt"
	hairColor      = "hcl"
	eyeColor       = "ecl"
	passportID     = "pid"

	mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
)

func main() {
	data := readFile()

	passports := parsePassport(data)

	validPassportsCount := validPassports(passports)

	fmt.Println("validPassports", validPassportsCount)
}

func validPassports(passports []string) int {
	var validPassports int
	for _, p := range passports {
		if checkMandatoryFields(p) {
			if validPassport(p) {
				validPassports++
			}
		}
	}
	return validPassports
}

func checkMandatoryFields(passport string) bool {
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
		//Split the fields here, for ex pid:021572410 as [pid, 021572410]
		//we need to check length because, there is space at the end we
		//are adding to while parsing passports
		if len(f) > 0 {
			fields := strings.Split(f, ":")
			field := fields[0]
			value := fields[1]
			switch field {
			case birthYear:
				y, _ := strconv.Atoi(value)
				valid := validate.BirthYear(y)
				if !valid {
					return false
				}
			case issueYear:
				y, _ := strconv.Atoi(value)
				valid := validate.IssueYear(y)
				if !valid {
					return false
				}
			case expirationYear:
				y, _ := strconv.Atoi(value)
				valid := validate.ExpirationYear(y)
				if !valid {
					return false
				}
			case height:
				valid := validate.Height(value)
				if !valid {
					return false
				}
			case hairColor:
				valid := validate.HairColor(value)
				if !valid {
					return false
				}
			case eyeColor:
				valid := validate.EyeColor(value)
				if !valid {
					return false
				}
			case passportID:
				valid := validate.PassportID(value)
				if !valid {
					return false
				}

			}
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

	file, _ := os.Open("./data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	input = append(input, "")
	return input
}
