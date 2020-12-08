package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	nums, err := readFile()
	if err != nil {
		fmt.Println("Err", err)
	}

	makeMap := convertSliceToDict(nums)

	num := findNumbersThatSum(makeMap, 2020)

	fmt.Println("num", num)
}

func readFile() ([]int, error) {

	file, err := os.Open("./data.txt")
	if err != nil {
		return nil, err
	}

	var perLine int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perLine)
		if err != nil {
			if err == io.EOF {
				return nums, nil
			}
			return nil, err
		}
		nums = append(nums, perLine)
	}
}

func convertSliceToDict(nums []int) map[int]int {
	m := make(map[int]int)

	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			m[num] = 1
		} else {
			m[num]++
		}
	}

	return m
}

func findNumbersThatSum(nums map[int]int, target int) int {
	for k := range nums {
		find := target - k
		_, ok := nums[find]
		if ok {
			return k * find
		}
	}
	return 0
}
