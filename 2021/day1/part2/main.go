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
		fmt.Println("Read input err: ", err)
		return
	}

	fmt.Println("Input: ", input)

	fmt.Println("Answer: ", getIncreaseCount(input))
}

func getIncreaseCount(nums []int) int {
	var count int

	for i := 0; i < len(nums)-3; i++ {
		num1 := nums[i] + nums[i+1] + nums[i+2]
		num2 := nums[i+1] + nums[i+2] + nums[i+3]

		if num2 > num1 {
			count++
		}
	}

	return count
}

func readInput() ([]int, error) {
	var output []int

	file, err := os.Open("./data.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		output = append(output, num)
	}

	return output, nil
}
