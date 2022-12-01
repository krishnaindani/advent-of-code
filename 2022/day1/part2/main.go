package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	data, err := readInput()
	if err != nil {
		fmt.Println("Read Input Error: ", err)
		return
	}

	maxCalories := elfWithThreeMostCalories(data)
	fmt.Println("maxCalories", maxCalories)
}

func elfWithThreeMostCalories(caloriesData map[int][]int) int {

	var calories []int

	for _, v := range caloriesData {
		calories = append(calories, sliceSum(v))
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] < calories[j]
	})
	fmt.Println("calories", calories)

	return calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]
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
