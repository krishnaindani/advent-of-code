package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := readInput()
	parsedData := parseInput(data)
	fmt.Println("parsedData", parsedData)
	allQuestionsAnswered := getQuestionsAnswered(parsedData)

	fmt.Println("Questions answered by all group:", allQuestionsAnswered)
}

func readInput() []string {

	file, err := os.Open("./sampleData.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var answers []string
	for scanner.Scan() {
		answers = append(answers, scanner.Text())
	}
	answers = append(answers, "")
	return answers
}

func parseInput(data []string) []string {

	var parsedData []string
	var text string

	for _, d := range data {
		if len(d) > 0 {
			text += d
		} else {
			parsedData = append(parsedData, text)
			text = ""
		}
	}

	return parsedData
}

func getQuestionsAnswered(questions []string) int {

	var questionsAnswered int

	for _, question := range questions {
		questionsAnswered += getOnePersonQuestionsAnswered(question)
	}

	return questionsAnswered
}

func getOnePersonQuestionsAnswered(question string) int {

	isPresent := make(map[int32]bool)
	var count int

	for _, char := range question {
		_, ok := isPresent[char]
		if !ok {
			count++
			isPresent[char] = true
		}
	}

	return count
}
