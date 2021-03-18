package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	makeMap, err := readInput()
	if err != nil {
		fmt.Println("Err", err)
	}

	value := findThreeNumbersThatSum(makeMap, 2020)

	fmt.Println("threeNumsMultiplication", value)

}

func readInput() (map[int]int, error) {

	m := make(map[int]int)
	file, err := os.Open("./data.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		_, ok := m[num]
		if !ok {
			m[num] = 1
		} else {
			m[num]++
		}
	}

	return m, nil
}

func findNumbersThatSum(nums map[int]int, target int) (bool, int, int) {
	for k := range nums {
		find := target - k
		_, ok := nums[find]
		if ok {
			return true, k, find
		}
	}
	return false, 0, 0
}

func findThreeNumbersThatSum(nums map[int]int, target int) int {
	for k := range nums {
		find := target - k
		present, num1, num2 := findNumbersThatSum(nums, find)
		if present {
			fmt.Printf("Num1: %d\nNum2: %d\nNum3: %d\n", k, num1, num2)
			return k * num1 * num2
		}
	}
	return 0
}
