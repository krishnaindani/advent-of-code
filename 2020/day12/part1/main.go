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

var (
	directions = []string{
		east,
		south,
		west,
		north,
	}

	directionIndex = map[string]int{
		east:  0,
		south: 1,
		west:  2,
		north: 3,
	}
)

func main() {
	input := readInput()

	fmt.Println("input", input)

	fmt.Println("answer", getManhattanCoordinates(input))
}

func move(x, y, value int, direction string) (int, int) {
	switch direction {
	case north:
		y += value
	case south:
		y -= value
	case east:
		x += value
	case west:
		x -= value
	}
	return x, y
}

func Abx(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getManhattanCoordinates(cd []coordinateData) int {

	var moveX, moveY int
	currentDirection := east

	for _, v := range cd {

		switch v.direction {
		case north:
			moveX, moveY = move(moveX, moveY, v.value, v.direction)
		case south:
			moveX, moveY = move(moveX, moveY, v.value, v.direction)
		case east:
			moveX, moveY = move(moveX, moveY, v.value, v.direction)
		case west:
			moveX, moveY = move(moveX, moveY, v.value, v.direction)
		case forward:
			moveX, moveY = move(moveX, moveY, v.value, currentDirection)
		case left:
			currentDirectionIndex := directionIndex[currentDirection]
			val := v.value / 90
			currentDirection = directions[(Abx(currentDirectionIndex-val))%4]
		case right:
			currentDirectionIndex := directionIndex[currentDirection]
			val := v.value / 90
			currentDirection = directions[(Abx(currentDirectionIndex+val))%4]
		}
		//fmt.Printf("x: %v, y: %v, direction: %v\n", moveX, moveY, currentDirection)
	}

	fmt.Printf("x: %v, y: %v, direction: %v\n", moveX, moveY, currentDirection)

	return Abx(moveX) + Abx(moveY)
}

func readInput() []coordinateData {

	f, _ := os.Open("./data.txt")

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
