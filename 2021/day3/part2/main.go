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

	oxygenGeneratorRatingBinary := getOxygenGeneratorRating(0, input)
	co2ScrubberRatingBinary := getCO2ScrubberRating(0, input)
	fmt.Println("oxygenGeneratorRatingBinary", oxygenGeneratorRatingBinary)
	fmt.Println("co2ScrubberRatingBinary", co2ScrubberRatingBinary)

	oxygenGeneratorRatingDecimal, err := strconv.ParseInt(oxygenGeneratorRatingBinary, 2, 64)
	if err != nil {
		fmt.Println("oxygenGeneratorRatingDecimal error: ", oxygenGeneratorRatingDecimal)
		return
	}

	co2ScrubberRatingDecimal, err := strconv.ParseInt(co2ScrubberRatingBinary, 2, 64)
	if err != nil {
		fmt.Println("co2ScrubberRatingDecimal error: ", co2ScrubberRatingDecimal)
		return
	}

	fmt.Println("oxygenGeneratorRatingDecimal: ", oxygenGeneratorRatingDecimal)
	fmt.Println("co2ScrubberRatingDecimal: ", co2ScrubberRatingDecimal)
	fmt.Println("Answer: ", oxygenGeneratorRatingDecimal*co2ScrubberRatingDecimal)
}

func getOxygenGeneratorRating(bit int, data []string) string {
	var output string

	dataLength := len(data)
	if dataLength == 1 || bit == len(data[0]) {
		return data[0]
	}

	var countOfZero, countOfOne int
	var countOfZeroElements, countOfOneElements []string
	for j := 0; j < dataLength; j++ {
		bit := string(data[j][bit])
		if bit == "0" {
			countOfZero++
			countOfZeroElements = append(countOfZeroElements, data[j])
		}
		if bit == "1" {
			countOfOne++
			countOfOneElements = append(countOfOneElements, data[j])
		}
	}
	if countOfZero > countOfOne {
		return getOxygenGeneratorRating(bit+1, countOfZeroElements)
	}
	if countOfOne > countOfZero {
		return getOxygenGeneratorRating(bit+1, countOfOneElements)
	}
	if countOfOne == countOfZero {
		return getOxygenGeneratorRating(bit+1, countOfOneElements)
	}
	return output
}

func getCO2ScrubberRating(bit int, data []string) string {
	var output string

	dataLength := len(data)
	if dataLength == 1 || bit == len(data[0]) {
		return data[0]
	}

	var countOfZero, countOfOne int
	var countOfZeroElements, countOfOneElements []string
	for j := 0; j < dataLength; j++ {
		bit := string(data[j][bit])
		if bit == "0" {
			countOfZero++
			countOfZeroElements = append(countOfZeroElements, data[j])
		}
		if bit == "1" {
			countOfOne++
			countOfOneElements = append(countOfOneElements, data[j])
		}
	}
	if countOfZero > countOfOne {
		return getCO2ScrubberRating(bit+1, countOfOneElements)
	}
	if countOfOne > countOfZero {
		return getCO2ScrubberRating(bit+1, countOfZeroElements)
	}
	if countOfOne == countOfZero {
		return getCO2ScrubberRating(bit+1, countOfZeroElements)
	}

	return output
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
