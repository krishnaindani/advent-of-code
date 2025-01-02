package main

import (
	"bufio"
	"fmt"
	"github/krishnaindani/advent-of-code/utils"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coordinates struct {
	x, y int
}

type ClawMachine struct {
	ButtonA Coordinates
	ButtonB Coordinates
	Prize   Coordinates
}

func main() {
	defer utils.TimeTrack(time.Now())
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("input", input)
	fmt.Println("tokensNeeded: ", calculateTokens(input))
}

func calculateTokens(clawMachines map[int]ClawMachine) int {
	tokens := 0

	for _, clawMachine := range clawMachines {
		tokens += calculateUsingCarmersRule(clawMachine)
	}

	return tokens
}

func calculateUsingCarmersRule(clawMachine ClawMachine) int {

	d := clawMachine.ButtonA.x*clawMachine.ButtonB.y - clawMachine.ButtonA.y*clawMachine.ButtonB.x
	d1 := clawMachine.Prize.x*clawMachine.ButtonB.y - clawMachine.Prize.y*clawMachine.ButtonB.x
	d2 := clawMachine.Prize.y*clawMachine.ButtonA.x - clawMachine.Prize.x*clawMachine.ButtonA.y

	if d1%d != 0 || d2%d != 0 {
		return 0
	}

	return (3 * (d1 / d)) + (d2 / d)
}

func readInput() (map[int]ClawMachine, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	out := make(map[int]ClawMachine)
	id := 0
	var clawMachine ClawMachine

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}
		button := strings.Split(line, ": ")[0]
		pos := strings.Split(line, ": ")[1]
		xStr, yStr := strings.Split(pos, ", ")[0], strings.Split(pos, ", ")[1]

		switch button {
		case "Button A":
			xNum := strings.Split(xStr, "X+")[1]
			yNum := strings.Split(yStr, "Y+")[1]
			xInt, _ := strconv.Atoi(xNum)
			yInt, _ := strconv.Atoi(yNum)
			clawMachine.ButtonA = Coordinates{xInt, yInt}
		case "Button B":
			xNum := strings.Split(xStr, "X+")[1]
			yNum := strings.Split(yStr, "Y+")[1]
			xInt, _ := strconv.Atoi(xNum)
			yInt, _ := strconv.Atoi(yNum)
			clawMachine.ButtonB = Coordinates{xInt, yInt}
		case "Prize":
			xNum := strings.Split(xStr, "X=")[1]
			yNum := strings.Split(yStr, "Y=")[1]
			xInt, _ := strconv.Atoi(xNum)
			yInt, _ := strconv.Atoi(yNum)
			clawMachine.Prize = Coordinates{xInt + 10000000000000, yInt + 10000000000000}
			out[id] = clawMachine
			id += 1

		}
	}

	return out, nil
}
