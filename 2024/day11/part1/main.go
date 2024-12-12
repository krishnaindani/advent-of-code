package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("input:", input)
	fmt.Println("stonesAfterAllBlink: ", stonesAfterAllBlink(input, 25))
}

func stonesAfterAllBlink(stones []int, freq int) int {

	stonesMap := intsSliceToMap(stones)

	for i := 0; i < freq; i++ {
		stonesMap = blinkAllStones(stonesMap)
	}

	return getSum(stonesMap)
}

func blinkAllStones(stones map[int]int) map[int]int {

	newStones := make(map[int]int)

	for stone, count := range stones {

		ns := blinkOnce(stone)

		for _, n := range ns {
			newStones[n] += count
		}
	}

	return newStones
}

func blinkOnce(stone int) []int {

	stoneStr := strconv.Itoa(stone)
	switch {
	case stone == 0:
		return []int{1}
	case len(stoneStr)%2 == 0:
		return splitEvenNumInTwo(stone)
	default:
		return []int{stone * 2024}

	}

	return nil
}

func splitEvenNumInTwo(num int) []int {

	numStr := strconv.Itoa(num)
	mid := len(numStr) / 2

	firstHalf := strings.TrimLeft(numStr[:mid], "0")
	secondHalf := strings.TrimLeft(numStr[mid:], "0")

	firstHalfNum, _ := strconv.Atoi(firstHalf)
	secondHalfNum, _ := strconv.Atoi(secondHalf)

	return []int{firstHalfNum, secondHalfNum}
}

func getSum(nums map[int]int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func readInput() ([]int, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Couldn't open file: %v\n", err)
		return nil, err
	}

	stones := make([]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		for _, v := range line {
			n, _ := strconv.Atoi(v)
			stones = append(stones, n)
		}
	}

	return stones, nil
}

func intsSliceToMap(s []int) map[int]int {
	m := make(map[int]int)
	for _, v := range s {
		m[v]++
	}
	return m
}
