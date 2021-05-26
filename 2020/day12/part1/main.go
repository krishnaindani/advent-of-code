package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coordinateData struct {
	direction string
	value     int
}

const (
	north   = "N"
	south   = "S"
	east    = "E"
	west    = "W"
	left    = "L"
	right   = "R"
	forward = "F"
)

func main() {
	input := readInput()

	fmt.Println("input", input)
}

func getManhattanCoordinates(cd []coordinateData) int {

	var moveX, moveY int

	for _, v := range cd {

		switch v.direction {
		case north:
			moveY += v.value
		case south:
			moveY -= v.value
		case east:
			moveX += v.value
		case west:
			moveX -= v.value
		case left:
		case right:
		case forward:
		}
	}

	return 0
}

func readInput() []coordinateData {

	f, _ := os.Open("./sampleData.txt")

	scanner := bufio.NewScanner(f)
	var output []coordinateData

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		value, _ := strconv.Atoi(line[1:])
		cd := coordinateData{
			direction: string(direction),
			value:     value,
		}
		output = append(output, cd)
	}

	return output
}
