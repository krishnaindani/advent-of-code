package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := readInput()

	fmt.Println(getNumberNotASumOfTwo(input, 25))
}

func getNumberNotASumOfTwo(nums []int, preamble int) int {

	preambleMap := make(map[int]bool)
	currentNumIndex := 0

	for i := 0; i < preamble; i++ {
		preambleMap[nums[i]] = true
	}

	for i := preamble; i < len(nums); i++ {
		if !isNumberSumOfTwo(preambleMap, nums[i]) {
			return findContiguousSubArraySum(nums, nums[i])
		}
		delete(preambleMap, nums[currentNumIndex])
		currentNumIndex++
		preambleMap[nums[i]] = true
	}

	return 0
}

func isNumberSumOfTwo(numsMap map[int]bool, num int) bool {

	for k, _ := range numsMap {
		numToFind := num - k
		if numToFind == k {
			continue
		}
		if _, ok := numsMap[numToFind]; ok {
			return true
		}
	}

	return false
}

func findContiguousSubArraySum(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		currSum := nums[i]
		currMin := nums[i]
		currMax := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if target == currSum {
				return currMin + currMax
			}

			if currSum > target {
				break
			}

			currMin = min(currMin, nums[j])
			currMax = max(currMax, nums[j])

			currSum += nums[j]

		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readInput() []int {
	var result []int

	f, _ := os.Open("./data.txt")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		result = append(result, v)
	}

	return result
}
