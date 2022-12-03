package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	A = "A"
	B = "B"
	C = "C"
	X = "X"
	Y = "Y"
	Z = "Z"

	RockPoints     = 1
	PaperPoints    = 2
	ScissorsPoints = 3

	LostPoints = 0
	DrawPoints = 3
	WinPoints  = 6

	Win  = "Win"
	Loss = "Loss"
	Draw = "Draw"

	Rock     = "Rock"
	Paper    = "Paper"
	Scissors = "Scissors"
)

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println("Read Input Error: ", err)
		return
	}

	fmt.Println("input:", input)
	fmt.Println("**************************************")
	fmt.Println("resultScore:", calculatePoints(input))
}

func calculatePoints(games [][]string) int {
	var resultScore int

	for _, game := range games {
		fmt.Println("game", game)
		opponent := getComponent(game[0])
		self := getComponent(game[1])
		fmt.Println(opponent, self)
		result := getResult(opponent, self)
		if result == Win {
			resultScore += WinPoints
		} else if result == Loss {
			resultScore += LostPoints
		} else if result == Draw {
			resultScore += DrawPoints
		}
		resultScore += getDefaultPoints(self)
		fmt.Println("resultScore", resultScore)
	}

	return resultScore
}

func getDefaultPoints(input string) int {
	switch input {
	case Rock:
		return RockPoints
	case Paper:
		return PaperPoints
	case Scissors:
		return ScissorsPoints
	}
	return 0
}

func getResult(a, b string) string {

	switch {
	case (a == Rock && b == Scissors) || (a == Scissors && b == Paper) || (a == Paper && b == Rock):
		return Loss
	case (a == Rock && b == Rock) || (a == Scissors && b == Scissors) || (a == Paper && b == Paper):
		return Draw
	}

	return Win
}

func getComponent(c string) string {
	switch {
	case c == X || c == A:
		return Rock
	case c == B || c == Y:
		return Paper
	case c == C || c == Z:
		return Scissors
	}
	return ""
}

func readInput() ([][]string, error) {
	var output [][]string

	file, err := os.Open("./data.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, strings.Split(line, " "))
	}

	return output, nil
}
