package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := readInput()
	if err != nil {
		fmt.Println("Read Input Error: ", err)
		return
	}

	maxCalories := elfWithMostCalories(data)
	fmt.Println("maxCalories", maxCalories)
}

func elfWithMostCalories(caloriesData map[int][]int) int {

	maxCalories := 0

	for _, v := range caloriesData {
		elfCaloriesSum := sliceSum(v)
		if elfCaloriesSum > maxCalories {
			maxCalories = elfCaloriesSum
		}
	}

	return maxCalories

}

func sliceSum(input []int) int {

	var total int

	for _, v := range input {
		total += v
	}

	return total
}

func readInput() (map[int][]int, error) {

	output := make(map[int][]int)

	file, err := os.Open("./data.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	currentElf := 1
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		if value == 0 {
			currentElf++
		} else {
			output[currentElf] = append(output[currentElf], value)
		}
	}

	return output, nil
}
