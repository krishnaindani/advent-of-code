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

	fmt.Println("computeCheckSum: ", computeCheckSum(input))
}

func computeCheckSum(nums []int) int {
	block := getBlock(nums)
	return getCheckSumResult(block)
}

func getCheckSumResult(nums []int) int {
	result := 0
	left := 0
	right := len(nums) - 1
	i := 0

	for left < right {
		if nums[left] == -1 {
			for nums[right] == -1 {
				right -= 1
			}
			nums[left], nums[right] = nums[right], nums[left]
		}

		result += i * nums[left]
		left += 1
		i++
	}

	return result
}

func getBlock(nums []int) []int {

	id := 0
	blocks := make([]int, 0)

	for i, num := range nums {
		if i%2 == 0 {
			blocks = append(blocks, buildBlock(id, num)...)
			id += 1
		} else {
			blocks = append(blocks, buildEmptyBlock(num)...)
		}
	}

	return blocks
}

func buildBlock(num int, frequency int) []int {

	block := make([]int, 0)

	for i := 0; i < frequency; i++ {
		block = append(block, num)
	}

	return block
}

func buildEmptyBlock(frequency int) []int {
	blocks := make([]int, 0)

	for i := 0; i < frequency; i++ {
		blocks = append(blocks, -1)
	}

	return blocks
}

func readInput() ([]int, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Couldn't open file: %v\n", err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Couldn't close file: %v\n", err)
		}
	}(file)

	nums := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for _, v := range line {
			i, _ := strconv.Atoi(v)
			nums = append(nums, i)
		}
	}

	return nums, nil
}
