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

	num1, num2 := findNumbersThatSum(makeMap, 2020)
	fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)
	fmt.Println("twoNumsMultiplication", num1*num2)

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

func findNumbersThatSum(nums map[int]int, target int) (int, int) {
	for k := range nums {
		find := target - k
		_, ok := nums[find]
		if ok {
			return k, find
		}
	}
	return 0, 0
}
