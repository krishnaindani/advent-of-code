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
	nums1, nums2, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	totalDistance := getSimilarityScore(nums1, nums2)
	fmt.Println("total distance", totalDistance)
}

func getSimilarityScore(nums1, nums2 map[int]int) int {
	score := 0

	for k, v := range nums1 {
		frequency, _ := nums2[k]
		score += k * v * frequency
	}

	return score
}

func readInput() (map[int]int, map[int]int, error) {

	first, second := make(map[int]int), make(map[int]int)

	file, err := os.Open("./data.txt")
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return nil, nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(value[0])
		b, _ := strconv.Atoi(value[len(value)-1])
		first[a] += 1
		second[b] += 1
	}

	return first, second, nil
}
