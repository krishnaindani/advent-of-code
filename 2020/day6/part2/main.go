package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := readInput()
	parsedData := parseInput(data)

	questionsAnswered := getQuestionsAnswered(parsedData)

	fmt.Println("No of questions answered:", questionsAnswered)

}

func readInput() []string {

	file, err := os.Open("./data.txt")
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

func parseInput(data []string) [][]string {

	var parsedData [][]string

	var groupQuestions []string
	for _, d := range data {
		if len(d) > 0 {
			groupQuestions = append(groupQuestions, d)
		} else {
			parsedData = append(parsedData, groupQuestions)
			groupQuestions = []string{}
		}
	}

	return parsedData
}

func getQuestionsAnswered(questions [][]string) int {

	var questionsAnswered int

	for _, question := range questions {
		questionsAnswered += getOneGroupQuestionsAnswered(question)
	}

	return questionsAnswered
}

func getOneGroupQuestionsAnswered(question []string) int {

	var count int
	length := len(question)

	if length == 0 {
		return count
	}

	if length == 1 {
		return len(question[0])
	}

	indexZero := question[0]

	for _, char := range indexZero {
		var stringsContains []bool
		for j := 1; j < length; j++ {
			stringsContains = append(stringsContains, strings.Contains(question[j], string(char)))
		}

		if getState(stringsContains) {
			count++
		}
		stringsContains = []bool{}
	}

	return count
}

func getState(states []bool) bool {
	state := states[0]
	for _, b := range states[1:] {
		state = state && b
	}
	return state
}
