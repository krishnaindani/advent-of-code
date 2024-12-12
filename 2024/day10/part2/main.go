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
	fmt.Println("getTrailHeadsScore:", getRatingsScore(input))
}

func getRatingsScore(trails [][]int) int {

	count := 0
	m := len(trails)
	n := len(trails[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if trails[i][j] == 0 {
				count += getRatingScore(Position{
					i, j,
				}, trails, make(map[Position]bool))
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

func getRatingScore(pos Position, trails [][]int, visited map[Position]bool) int {

	if trails[pos.x][pos.y] == 9 {
		return 1
	}

	count := 0
	visited[pos] = true
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for _, direction := range directions {
		dx := direction[0] + pos.x
		dy := direction[1] + pos.y
		if IsCoordinatesValid(dx, dy, trails) &&
			!visited[Position{dx, dy}] &&
			trails[dx][dy] == trails[pos.x][pos.y]+1 {
			count += getRatingScore(Position{dx, dy}, trails, visited)
		}
	}

	visited[pos] = false
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
