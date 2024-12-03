package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result", getResult(input))
}

func getResult(nums [][]int) int {
	answer := 0

	for _, num := range nums {
		answer += num[0] * num[1]
	}

	return answer
}

func readInput() ([][]int, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		log.Printf("error opening file: %v", err)
		return nil, err
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}(file)

	answer := make([][]int, 0)
	re := regexp.MustCompile("mul\\(\\d+,\\d+\\)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			strs := regexp.MustCompile("\\d+").FindAllString(match, -1)
			num := make([]int, 0, 2)
			for _, str := range strs {
				i, _ := strconv.Atoi(str)
				num = append(num, i)
			}
			answer = append(answer, num)
		}
	}

	return answer, nil
}
