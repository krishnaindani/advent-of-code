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

	fmt.Println("one and three jolts multiplication", calculateJoltDifferenceMultiplication(input))
}

func calculateJoltDifferenceMultiplication(jolts []int) int {

	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	frequencyStore := map[int]int{}

	for i, v := range jolts {
		if i == 0 {
			frequencyStore[v-0]++
		} else {
			frequencyStore[v-jolts[i-1]]++
		}
	}

	return frequencyStore[3] * frequencyStore[1]
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
