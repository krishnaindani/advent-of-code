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

	fmt.Println("calculateTotalFencingCost: ", calculateTotalFencingCost(input))
}

func calculateTotalFencingCost(grid [][]string) int {

	visited := make(map[Position]bool)
	cost := 0
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[Position{i, j}] {
				visited[Position{i, j}] = true
				cost += getPlantGardenCost(Position{i, j}, grid, visited)
			}
		}
	}

	return cost
}

func getPlantGardenCost(pos Position,
	grid [][]string,
	visited map[Position]bool) int {

	area := 0
	perimeter := 0
	plant := grid[pos.x][pos.y]
	queue := make([][]int, 0)
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	queue = append(queue, []int{pos.x, pos.y})

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y := curr[0], curr[1]
		area += 1

		for _, direction := range directions {
			dx := direction[0] + x
			dy := direction[1] + y
			if isCoordinatesValid(dx, dy, grid, plant) {
				if !visited[Position{dx, dy}] {
					visited[Position{dx, dy}] = true
					queue = append(queue, []int{dx, dy})
				}
			} else {
				perimeter += 1
			}
		}

	}

	return area * perimeter
}

func isCoordinatesValid(x, y int, grid [][]string, plant string) bool {
	m := len(grid)
	n := len(grid[0])
	return (x >= 0 && x < m) && (y >= 0 && y < n) && grid[x][y] == plant
}

func readInput() ([][]string, error) {

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

	grid := make([][]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
	}

	return grid, nil
}
