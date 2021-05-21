package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := readInput()
	fmt.Println("input", input)

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
			return nums[i]
		}
		delete(preambleMap, nums[currentNumIndex])
		currentNumIndex++
		preambleMap[nums[i]] = true
	}

	return 0
}

func isNumberSumOfTwo(numsMap map[int]bool, num int) bool {

	for k, _ := range numsMap {
		fmt.Println("num", num, "numMap", numsMap)
		numToFind := num - k
		if numToFind == k {
			continue
		}
		fmt.Println("k", k, "numToFind", numToFind)
		if _, ok := numsMap[numToFind]; ok {
			return true
		}
	}

	return false
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
