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
	rules, pages, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("score: ", getScore(pages, rules))
}

func getScore(pages, rules [][]int) int {
	score := 0

	for _, page := range pages {
		if isOrdered(rules, page) {
			score += page[len(page)/2]
		}
	}

	return score
}

func isOrdered(rules [][]int, page []int) bool {

	numToIndex := make(map[int]int)

	for i, p := range page {
		numToIndex[p] = i
	}

	for _, rule := range rules {
		x, y := rule[0], rule[1]
		v1, exists := numToIndex[x]
		v2, exists2 := numToIndex[y]
		if exists && exists2 && (v1 > v2) {
			return false
		}
	}

	return true
}

func readInput() ([][]int, [][]int, error) {

	rules := make([][]int, 0)
	pages := make([][]int, 0)

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return nil, nil, err
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Printf("error closing file: %v\n", err)
		}
	}(file)

	isRule := true
	isPage := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRule = !isRule
			isPage = !isPage
			continue
		}

		if isRule {
			element := strings.Split(line, "|")
			num1, _ := strconv.Atoi(element[0])
			nums2, _ := strconv.Atoi(element[1])
			rules = append(rules, []int{num1, nums2})
		}

		if isPage {
			page := make([]int, 0)
			elements := strings.Split(line, ",")
			for _, element := range elements {
				num, _ := strconv.Atoi(element)
				page = append(page, num)
			}
			pages = append(pages, page)
		}
	}

	return rules, pages, nil
}
