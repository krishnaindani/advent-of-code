package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	x, y int
}

type Robot struct {
	pos Coordinates
	vel Coordinates
}

func main() {

	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("input", input)
	fmt.Println("score: ", calculateSafetyScore(input, Coordinates{101, 103}))
}

func calculateSafetyScore(robots []Robot, bounds Coordinates) int {

	for i := 0; i < 100; i++ {
		moveRobots(robots, bounds)
	}

	xMid := bounds.x / 2
	yMid := bounds.y / 2
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, robot := range robots {

		if robot.pos.x == xMid && robot.pos.y == xMid {
			continue
		}

		//first quadrant
		if robot.pos.x < xMid && robot.pos.y < yMid {
			q1 += 1
		} else if robot.pos.x > xMid && robot.pos.y < yMid {
			//second quadrant
			q2 += 1
		} else if robot.pos.x < xMid && robot.pos.y > yMid {
			//third quadrant
			q3 += 1
		} else if robot.pos.x > xMid && robot.pos.y > yMid {
			//fourth quadrant
			q4 += 1
		}
	}

	return q1 * q2 * q3 * q4
}

func moveRobots(robots []Robot, bounds Coordinates) {

	for i := range robots {
		robots[i].pos.x += robots[i].vel.x
		robots[i].pos.y += robots[i].vel.y

		robots[i].pos.x = ((robots[i].pos.x % bounds.x) + bounds.x) % bounds.x
		robots[i].pos.y = ((robots[i].pos.y % bounds.y) + bounds.y) % bounds.y
	}

}

func readInput() ([]Robot, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var robot []Robot

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		p, v := line[0], line[1]
		pStr := strings.Split(strings.Split(p, "p=")[1], ",")
		vStr := strings.Split(strings.Split(v, "v=")[1], ",")
		px, _ := strconv.Atoi(pStr[0])
		py, _ := strconv.Atoi(pStr[1])
		vx, _ := strconv.Atoi(vStr[0])
		vy, _ := strconv.Atoi(vStr[1])
		robot = append(robot, Robot{
			pos: Coordinates{
				x: px,
				y: py,
			},
			vel: Coordinates{
				x: vx,
				y: vy,
			},
		})

	}

	return robot, nil
}
