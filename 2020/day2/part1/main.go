package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile()

	fmt.Println("Number of valid passwords are:", validPasswords(input))
}

func validPasswords(data []string) int {
	var validPasswords int
	for _, v := range data {
		isValid := validPassword(v)
		if isValid {
			validPasswords++
		}
	}

	return validPasswords
}

func validPassword(password string) bool {

	s := strings.Fields(password)

	digitsSlice := strings.Split(s[0], "-")
	minDigits, _ := strconv.Atoi(digitsSlice[0])
	maxDigits, _ := strconv.Atoi(digitsSlice[1])
	alphabetSupported := strings.Split(s[1], ":")[0]
	passwordString := s[2]

	letterCount := strings.Count(passwordString, alphabetSupported)
	if letterCount >= minDigits && letterCount <= maxDigits {
		return true
	}

	return false
}

func readFile() []string {
	var input []string

	file, _ := os.Open("./data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
