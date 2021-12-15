package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	forward = "forward"
	down    = "down"
	up      = "up"
)

func main() {
	directions, angles, err := readInput()
	if err != nil {
		fmt.Println("Read Input Error: ", err)
		return
	}

	fmt.Println("Directions: ", directions)
	fmt.Println("Angles: ", angles)

	direction, angle := getDirection(directions, angles)
	fmt.Println("direction: ", direction)
	fmt.Println("angle: ", angle)
	fmt.Println("Answer: ", direction*angle)
}

func getDirection(directions []string, angle []int) (int, int) {
	var d, a int

	for i, v := range directions {
		if v == forward {
			d += angle[i]
		}
		if v == down {
			a += angle[i]
		}
		if v == up {
			a -= angle[i]
		}
	}

	return d, a
}

func readInput() ([]string, []int, error) {

	var directions []string
	var angle []int

	file, err := os.Open("./dcata.txt")
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		directions = append(directions, s[0])
		num, _ := strconv.Atoi(s[1])
		angle = append(angle, num)
	}

	return directions, angle, nil
}
