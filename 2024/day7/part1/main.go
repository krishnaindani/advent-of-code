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
	
	fmt.Println("getCalibrationResult: ", getCalibrationResult(input))
}

func getCalibrationResult(data map[int][]int) int {

	result := 0

	for target, nums := range data {
		if canFormTarget(target, nums) {
			result += target
		}
	}

	return result
}

func canFormTarget(target int, nums []int) bool {

	n := len(nums)
	var dfs func(index int, current int) bool

	dfs = func(index int, current int) bool {

		if index == n {
			return current == target
		}

		return dfs(index+1, current*nums[index]) ||
			dfs(index+1, current+nums[index])
	}

	return dfs(1, nums[0])
}

func readInput() (map[int][]int, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Couldn't open file: %v\n", err)
		return nil, err
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Printf("Couldn't close file: %v\n", err)
		}
	}(file)

	store := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		target, _ := strconv.Atoi(strings.TrimSpace(line[0]))

		numStrs := strings.Fields(line[1])
		nums := make([]int, 0, len(numStrs))

		for _, str := range numStrs {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}

		store[target] = nums
	}

	return store, nil
}
