package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Antenna1 Position (x1,y1)
//Antenna2 Position (x2,y2)
//Find point p where
//1. All three points are collinear using slop formula

type Position struct {
	x, y int
}

func main() {

	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("countAntinodes: ", countAntinodes(input))
}

func countAntinodes(grid [][]rune) int {

	frequencies := getFrequencies(grid)
	allAntinodes := make(map[Position]bool)

	for _, posn := range frequencies {
		if len(posn) < 2 {
			continue
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				p := Position{i, j}
				for a := 0; a < len(posn); a++ {
					for b := a + 1; b < len(posn); b++ {
						if isCollinear(p, posn[a], posn[b]) {
							allAntinodes[p] = true
							break
						}
					}
				}
			}
		}
	}

	return len(allAntinodes)
}

func getFrequencies(grid [][]rune) map[rune][]Position {

	frequencies := make(map[rune][]Position)
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			char := grid[i][j]
			if char != '.' {
				frequencies[char] = append(frequencies[char], Position{i, j})
			}
		}
	}

	return frequencies
}

func isCollinear(p, a, b Position) bool {

	dx1 := a.x - p.x
	dy1 := a.y - p.y
	dx2 := b.x - p.x
	dy2 := b.y - p.y

	return dx1*dy2 == dx2*dy1
}

func readInput() ([][]rune, error) {

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

	grid := make([][]rune, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid, nil
}
