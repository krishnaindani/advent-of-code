package validate

import (
	"regexp"
	"strconv"
	"strings"
)

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
