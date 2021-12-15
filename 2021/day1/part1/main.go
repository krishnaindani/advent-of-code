package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println("Read Input Error: ", err)
		return
	}

	fmt.Println("Input: ", input)

	fmt.Println("Answer: ", increaseCount(input))
}

func increaseCount(nums []int) int {
	var count int

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		}
	}

	return count
}

func readInput() ([]int, error) {

	var data []int

	file, err := os.Open("./data.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		data = append(data, num)
	}

	return data, nil
}
