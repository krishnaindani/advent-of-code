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

	for !instructions[index].visited {

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
			continue
		}
	}

	return accumulator
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
			visited:   false,
		})
	}

	return result
}
