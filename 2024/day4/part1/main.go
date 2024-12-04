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

	isValid := func(x, y int) bool {
		return (x >= 0 && x < m) && (y >= 0 && y < n)
	}

	searchFromPosition := func(x, y, dx, dy int) bool {
		for i := 0; i < len(word); i++ {
			nx := x + i*dx
			ny := y + i*dy
			if !isValid(nx, ny) || lines[nx][ny] != word[i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, direction := range directions {
				if searchFromPosition(i, j, direction[0], direction[1]) {
					count += 1
				}
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
