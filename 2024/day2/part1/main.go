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

	totalSafeReports := getSafeReports(input)
	fmt.Println("total safe reports", totalSafeReports)

}

func getSafeReports(input [][]int) int {

	count := 0

	for _, row := range input {
		if (isSortedAscending(row) || isSortedDescending(row)) && checkDiff(row) {
			count += 1
		}
	}

	return count
}

func isSortedAscending(nums []int) bool {
	last := nums[0]

	for i := 1; i < len(nums); i++ {
		if last >= nums[i] {
			return false
		}
		last = nums[i]
	}
	return true
}

func isSortedDescending(nums []int) bool {
	last := nums[0]

	for i := 1; i < len(nums); i++ {
		if last <= nums[i] {
			return false
		}
		last = nums[i]
	}
	return true
}

func checkDiff(nums []int) bool {
	last := nums[0]

	for i := 1; i < len(nums); i++ {
		v := abs(nums[i] - last)
		if !(v >= 1 && v <= 3) {
			return false
		}
		last = nums[i]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readInput() ([][]int, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	numbers := make([][]int, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		num := make([]int, 0)
		for _, v := range line {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			num = append(num, vInt)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
