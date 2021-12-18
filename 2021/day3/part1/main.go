package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println("Read input error: ", err)
		return
	}

	fmt.Println("Input: ", input)

	gammaRateBinary, epsilonRateBinary := getGammaAndEpsilonRate(input)
	fmt.Println("gammaRateBinary: ", gammaRateBinary)
	fmt.Println("epsilonRateBinary: ", epsilonRateBinary)

	gammaRateDecimal, err := strconv.ParseInt(gammaRateBinary, 2, 64)
	if err != nil {
		fmt.Println("gammaRateDecimal error: ", gammaRateDecimal)
		return
	}

	epsilonRateDecimal, err := strconv.ParseInt(epsilonRateBinary, 2, 64)
	if err != nil {
		fmt.Println("epsilonRateDecimal error: ", epsilonRateDecimal)
		return
	}

	fmt.Println("gammaRateDecimal: ", gammaRateDecimal)
	fmt.Println("epsilonRateDecimal: ", epsilonRateDecimal)
	fmt.Println("Answer: ", gammaRateDecimal*epsilonRateDecimal)
}

func getGammaAndEpsilonRate(data []string) (string, string) {
	var gammaRateBinary, epsilonRateBinary string

	bitCount := len(data[0])

	for i := 0; i < bitCount; i++ {

		var countOfZero, countOfOne int

		for j := 0; j < len(data); j++ {
			bit := string(data[j][i])
			if bit == "0" {
				countOfZero++
			}
			if bit == "1" {
				countOfOne++
			}
		}

		if countOfZero > countOfOne {
			gammaRateBinary += "0"
			epsilonRateBinary += "1"
		}
		if countOfOne > countOfZero {
			gammaRateBinary += "1"
			epsilonRateBinary += "0"
		}
	}
	return gammaRateBinary, epsilonRateBinary
}

func readInput() ([]string, error) {

	file, err := os.Open("./data.txt")
	if err != nil {
		return nil, err
	}

	var output []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output, nil
}
