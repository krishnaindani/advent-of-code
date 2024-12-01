package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums1, nums2, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	totalDistance := getTotalDistance(nums1, nums2)
	fmt.Println("total distance", totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getTotalDistance(nums1, nums2 []int) int {
	distance := 0

	for i := range nums1 {
		distance += abs(nums1[i] - nums2[i])
	}

	return distance
}

func readInput() ([]int, []int, error) {

	first, second := make([]int, 0), make([]int, 0)

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
		first = append(first, a)
		second = append(second, b)
	}

	return first, second, nil
}
