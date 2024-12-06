package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func getScore(pages [][]int, rules map[int]map[int]struct{}) int {
	score := 0

	for _, page := range pages {
		if !isOrdered(rules, page) {
			newPage := fixOrder(rules, page)
			score += newPage[len(newPage)/2]
		}
	}

	return score
}

func fixOrder(rules map[int]map[int]struct{}, page []int) []int {

	slices.SortFunc(page, func(a, b int) int {

		if _, ok := rules[a][b]; ok {
			return -1
		}

		if _, ok := rules[b][a]; !ok {
			return 1
		}

		return 0
	})

	return page
}

func isOrdered(rules map[int]map[int]struct{}, page []int) bool {

	n := len(page)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			char1 := page[i]
			char2 := page[j]
			if _, ok := rules[char2][char1]; ok {
				return false
			}
		}
	}

	return true
}

func readInput() (map[int]map[int]struct{}, [][]int, error) {

	rulesMap := make(map[int]map[int]struct{})
	pages := make([][]int, 0)

	file, err := os.Open("sampleData.txt")
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
			if _, ok := rulesMap[num1]; !ok {
				rulesMap[num1] = make(map[int]struct{})
			}
			rulesMap[num1][nums2] = struct{}{}
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

	return rulesMap, pages, nil
}
