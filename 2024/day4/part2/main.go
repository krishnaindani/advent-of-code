package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	directions = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}
	word = "XMAS"
)

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count", getOccurenceCount(input))
}

func getOccurenceCount(lines []string) int {

	count := 0
	m := len(lines)
	n := len(lines[0])

	searchFromPosition := func(x, y int) bool {

		/*
			M.S
			.A.
			M.S
		*/
		forward := lines[x][y] == 'M' &&
			lines[x][y+2] == 'S' &&
			lines[x+1][y+1] == 'A' &&
			lines[x+2][y] == 'M' &&
			lines[x+2][y+2] == 'S'

		/*
			S.S
			.A.
			M.M
		*/
		upward := lines[x][y] == 'S' &&
			lines[x][y+2] == 'S' &&
			lines[x+1][y+1] == 'A' &&
			lines[x+2][y] == 'M' &&
			lines[x+2][y+2] == 'M'

		/*
			M.M
			.A.
			S.S
		*/

		diagonalOne := lines[x][y] == 'M' &&
			lines[x][y+2] == 'M' &&
			lines[x+1][y+1] == 'A' &&
			lines[x+2][y] == 'S' &&
			lines[x+2][y+2] == 'S'

		/*
			S.M
			.A.
			S.M
		*/
		diagonalTwo := lines[x][y] == 'S' &&
			lines[x][y+2] == 'M' &&
			lines[x+1][y+1] == 'A' &&
			lines[x+2][y] == 'S' &&
			lines[x+2][y+2] == 'M'

		return forward || upward || diagonalOne || diagonalTwo
	}

	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			if searchFromPosition(i, j) {
				count += 1
			}
		}
	}

	return count
}

func readInput() ([]string, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return nil, err
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v\n", err)
		}
	}(file)

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
