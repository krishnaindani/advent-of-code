package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	jump        = "jmp"
	accumulate  = "acc"
	noOperation = "nop"
)

type instruction struct {
	operation string
	value     int
}

func main() {
	input := readInput()
	fmt.Println("input", input)

	fmt.Println("accumulator", getValueBeforeInfiniteLoopStarts(input))
}

func getValueBeforeInfiniteLoopStarts(instructions []instruction) int {

	for i := range instructions {

		//create the copy each time
		instructionsCopy := make([]instruction, len(instructions))
		copy(instructionsCopy, instructions)

		switch instructions[i].operation {
		case jump:
			instructionsCopy[i].operation = noOperation
		case noOperation:
			instructionsCopy[i].operation = jump
		case accumulate:
			continue
		}

		//will check if it is infinite loop or not
		if v, ok := isInfiniteLoop(instructionsCopy); !ok {
			return v
		}

	}

	return -1
}

func isInfiniteLoop(instructions []instruction) (int, bool) {

	visited := map[int]bool{}
	var accumulator int
	index := 0

	for index < len(instructions) {

		if visited[index] {
			return 0, true
		}

		switch instructions[index].operation {
		case accumulate:
			accumulator += instructions[index].value
			visited[index] = true
			index++
		case jump:
			visited[index] = true
			index += instructions[index].value
		case noOperation:
			visited[index] = true
			index++
		}
	}

	return accumulator, false
}

func readInput() []instruction {
	f, _ := os.Open("./data.txt")

	var result []instruction
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		arg, _ := strconv.Atoi(line[1])
		result = append(result, instruction{
			operation: line[0],
			value:     arg,
		})
	}

	return result
}
