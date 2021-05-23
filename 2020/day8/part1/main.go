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
	var accumulator int

	index := 0

	visited := map[int]bool{}

	for !visited[index] {

		//is already visited
		if _, ok := visited[index]; ok {
			return accumulator
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
		})
	}

	return result
}
