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

	winningMappings := map[string]int{
		"X": LostPoints,
		"Y": DrawPoints,
		"Z": WinPoints,
	}

	for _, game := range games {
		resultScore += winningMappings[game[1]]

		//switch on opponent side
		switch game[0] {

		//opp: Rock
		case A:
			switch game[1] {
			//loss
			case X:
				resultScore += ScissorsPoints
				//draw
			case Y:
				resultScore += RockPoints
				//win
			case Z:
				resultScore += PaperPoints
			}

			//opp: Paper
		case B:
			switch game[1] {
			//loss
			case X:
				resultScore += RockPoints
				//draw
			case Y:
				resultScore += PaperPoints
				//win
			case Z:
				resultScore += ScissorsPoints
			}

			//opp: Scissors
		case C:
			switch game[1] {
			//loss
			case X:
				resultScore += PaperPoints
				//draw
			case Y:
				resultScore += ScissorsPoints
				//win
			case Z:
				resultScore += RockPoints
			}
		}

	}

	return resultScore
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
