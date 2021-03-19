package validate

import (
	"regexp"
	"strconv"
	"strings"
)

//BirthYear ...
func BirthYear(birthYear int) bool {
	return (birthYear >= 1920) && (birthYear <= 2002)
}

//IssueYear ...
func IssueYear(issueYear int) bool {
	return (issueYear >= 2010) && (issueYear <= 2020)
}

//ExpirationYear ...
func ExpirationYear(expirationYear int) bool {
	return (expirationYear >= 2020) && (expirationYear <= 2030)
}

//EyeColor ...
func EyeColor(color string) bool {

	return (color == "amb") || (color == "blu") || (color == "gry") ||
		(color == "brn") || (color == "hzl") || (color == "oth") ||
		(color == "grn")
}

//PassportID ...
func PassportID(id string) bool {
	return len(id) == 9
}

//HairColor ...
func HairColor(color string) bool {
	match, _ := regexp.MatchString("^#([0-9a-f]{6})", color)
	return match
}

//Height ...
func Height(height string) bool {

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
