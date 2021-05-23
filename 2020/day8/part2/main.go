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
	visited   bool
}

func main() {
	input := readInput()
	fmt.Println("input", input)

	fmt.Println("accumulator", getValueBeforeInfiniteLoopStarts(input))
}

func getValueBeforeInfiniteLoopStarts(instructions []instruction) int {
	var accumulator int

	index := 0

	for !instructions[index].visited && index < len(instructions) {
		fmt.Println("index", index, "instructions", instructions, "accumulator", accumulator)
		switch instructions[index].operation {
		case accumulate:
			accumulator += instructions[index].value
			instructions[index].visited = true
			index++
		case jump:
			v, ok := canBreakInfiniteLoop(instructions, index, jump, accumulator)
			if ok {
				return v
			}
			instructions[index].visited = true
			index += instructions[index].value
		case noOperation:
			v, ok := canBreakInfiniteLoop(instructions, index, noOperation, accumulator)
			if ok {
				return v
			}
			instructions[index].visited = true
			index++
			continue
		}
	}

	return accumulator
}

func canBreakInfiniteLoop(instructionsCopy []instruction,
	index int, operationType string, accumulator int) (int, bool) {

	instructions := make([]instruction, len(instructionsCopy))
	copy(instructions, instructionsCopy)

	if instructions[index].operation == jump {
		instructions[index].operation = noOperation
	} else if instructions[index].operation == noOperation {
		instructions[index].operation = jump
	}

	for index < len(instructions) {
		fmt.Println("from loop index", index, "instructions", instructions, "accumulator", accumulator)
		if !instructions[index].visited {
			return 0, false
		}
		switch instructions[index].operation {
		case accumulate:
			accumulator += instructions[index].value
			instructions[index].visited = true
			index++
		case jump:
			instructions[index].visited = true
			index += instructions[index].value
		case noOperation:
			instructions[index].visited = true
			index++
		}
	}

	if index == len(instructions) {
		return accumulator, true
	}

	return 0, false
}

func readInput() []instruction {
	f, _ := os.Open("./sampleData.txt")

	var result []instruction
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		arg, _ := strconv.Atoi(line[1])
		result = append(result, instruction{
			operation: line[0],
			value:     arg,
			visited:   false,
		})
	}

	return result
}
