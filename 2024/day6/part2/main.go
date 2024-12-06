package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	fmt.Println("getPath", getPath(input))
}

func getPath(grid [][]string) int {

	queue := make([][]int, 0)
	visited := make(map[Position]bool)
	m := len(grid)
	n := len(grid[0])

	isValid := func(x int, y int) bool {
		return (x >= 0 && x < m) && (y >= 0 && y < n)
	}

top:
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == "^" {
				queue = append(queue, []int{i, j})
				grid[i][j] = "."
				break top
			}
		}
	}

	directionToCoordinates := map[int][]int{
		0: []int{-1, 0},
		1: []int{0, 1},
		2: []int{1, 0},
		3: []int{0, -1},
	}

	currDirection := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y := curr[0], curr[1]
		visited[Position{x, y}] = true
		dx := x + directionToCoordinates[currDirection][0]
		dy := y + directionToCoordinates[currDirection][1]
		if isValid(dx, dy) {
			if grid[dx][dy] == "." {
				queue = append(queue, []int{dx, dy})
			}
			if grid[dx][dy] == "#" {
				currDirection = (currDirection + 1) % 4
				queue = append(queue, []int{
					x + directionToCoordinates[currDirection][0],
					y + directionToCoordinates[currDirection][1],
				})
			}
		} else {
			break
		}
	}

	return len(visited)
}

func readInput() ([][]string, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	grid := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	return grid, nil
}
