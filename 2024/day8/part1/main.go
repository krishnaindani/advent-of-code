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
//2. Distance is twice
//distance(p, antenna1) = 2 * distance(p, antenna2)
//distance(p, antenna2) = 2 * distance(p, antenna1)

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
		for i := 0; i < len(posn); i++ {
			for j := i + 1; j < len(posn); j++ {
				antinodes := findAntinodes(grid, posn[i], posn[j])
				for _, antinode := range antinodes {
					if isInBounds(antinode, grid) {
						allAntinodes[antinode] = true
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

func findAntinodes(grid [][]rune, a, b Position) []Position {

	antinodes := make([]Position, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isAntinode(Position{i, j}, a, b) {
				antinodes = append(antinodes, Position{i, j})
			}
		}
	}

	return antinodes
}

func isAntinode(p, a, b Position) bool {

	if !isCollinear(p, a, b) {
		return false
	}

	dap := getDistanceSquared(a, p)
	dbp := getDistanceSquared(b, p)

	return dap == 4*dbp || dbp == 4*dap
}

func isInBounds(p Position, grid [][]rune) bool {
	m := len(grid)
	n := len(grid[0])
	return (p.x >= 0 && p.x < m) && (p.y >= 0 && p.y < n)
}

// squared distance to avoid negative
func getDistanceSquared(a, b Position) int {
	x := a.x - b.x
	y := a.y - b.y
	return x*x + y*y
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
