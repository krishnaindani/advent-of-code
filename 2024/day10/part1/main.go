package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("input:", input)
	fmt.Println("getTrailHeadsScore:", getTrailHeadsScore(input))
}

func getTrailHeadsScore(trails [][]int) int {

	count := 0
	m := len(trails)
	n := len(trails[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if trails[i][j] == 0 {
				count += getTrailHeadScore(Position{
					i, j,
				}, trails)
			}
		}
	}

	return count
}

func IsCoordinatesValid(x, y int, trails [][]int) bool {
	m := len(trails)
	n := len(trails[0])
	return (x >= 0 && x < m) && (y >= 0 && y < n)
}

func getTrailHeadScore(start Position, trails [][]int) int {

	count := 0
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	queue := make([][]int, 0)
	visited := make(map[Position]bool)

	queue = append(queue, []int{start.x, start.y, 0})
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y, prevVal := curr[0], curr[1], curr[2]
		if prevVal == 9 {
			count += 1
			continue
		}

		for _, direction := range directions {
			dx := x + direction[0]
			dy := y + direction[1]
			if IsCoordinatesValid(dx, dy, trails) &&
				!visited[Position{dx, dy}] &&
				trails[dx][dy] == prevVal+1 {
				queue = append(queue, []int{dx, dy, trails[dx][dy]})
				visited[Position{dx, dy}] = true
			}
		}
	}

	return count
}

func readInput() ([][]int, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Couldn't open file: %v\n", err)
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Couldn't close file: %v\n", err)
		}
	}()

	matrix := make([][]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		nums := make([]int, len(line))
		for i, v := range line {
			nums[i], _ = strconv.Atoi(v)
		}
		matrix = append(matrix, nums)
	}

	return matrix, nil
}
