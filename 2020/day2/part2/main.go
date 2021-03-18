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

	fmt.Println("Number of valid passwords:", validPasswords(input))
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

	minWord := string(passwordString[minDigits-1]) == alphabetSupported
	maxWord := string(passwordString[maxDigits-1]) == alphabetSupported

	if minWord != maxWord == true {
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
