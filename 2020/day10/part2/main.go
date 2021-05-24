package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := readInput()

	fmt.Println("input", input)

	fmt.Println("different combination:", calculateJoltDifferenceMultiplication(input))
}

//Example
//1 4 5 6 7 10 11 12 15 16 19 22
//
//1 4 5 7 10 11 12 15 16 19 22
//1 4 6 7 10 11 12 15 16 19 22
//1 4 7 10 11 12 15 16 19 22
//
//1 4 5 7 10 12  15 16 19 22
//1 4 6 7 10  12 15 16 19 22
//1 4 7 10  12 15 16 19 22
func calculateJoltDifferenceMultiplication(jolts []int) int {

	//adding for the first step
	jolts = append(jolts, 0)

	sort.Ints(jolts)

	//adding the last one
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	dp := make([]int, len(jolts))
	dp[0] = 1

	for i := 1; i < len(jolts); i++ {
		currJolt := jolts[i]
		for j := i - 1; j >= 0; j-- {
			if currJolt-jolts[j] <= 3 {
				dp[i] += dp[j]
			} else {
				break
			}
		}
	}

	fmt.Println("dp", dp)

	return dp[len(jolts)-1]
}

func readInput() []int {
	var jolts []int

	f, _ := os.Open("./data.txt")

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		jolt, _ := strconv.Atoi(scanner.Text())
		jolts = append(jolts, jolt)
	}

	return jolts
}
