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

	totalSafeReports := getSafeReports(input)
	fmt.Println("total safe reports", totalSafeReports)

}

func getSafeReports(input [][]int) int {

	count := 0

	for _, row := range input {
		if isReportSafe(row) || checkReportSafetyWithDeletion(row) {
			count += 1
		}
	}

	return count
}

func isReportSafe(report []int) bool {

	flagIncrease, flagDecrease := false, false

	for i := 1; i < len(report); i++ {
		v := report[i] - report[i-1]

		if v > 0 {
			flagIncrease = true
		} else if v < 0 {
			flagDecrease = true
		} else {
			return false
		}

		if flagIncrease && flagDecrease {
			return false
		}

		if v > 3 || v < -3 {
			return false
		}

	}

	return true
}

func checkReportSafetyWithDeletion(report []int) bool {

	for i := 0; i < len(report); i++ {
		flag := checkReportSafetyWithDeletionIndex(report, i)
		if flag {
			return true
		}

	}

	return false
}

func checkReportSafetyWithDeletionIndex(report []int, index int) bool {
	newReport := make([]int, len(report))
	copy(newReport, report)

	if index == len(report)-1 {
		newReport = newReport[:index]
	} else {
		newReport = append(newReport[:index], newReport[index+1:]...)
	}

	return isReportSafe(newReport)
}

func readInput() ([][]int, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	numbers := make([][]int, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		num := make([]int, 0)
		for _, v := range line {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			num = append(num, vInt)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
