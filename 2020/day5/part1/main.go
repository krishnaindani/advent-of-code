package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	front = "f"
	back  = "b"

	right = "r"
	left  = "l"
)

func main() {
	seats := readInput()

	var best int
	//first calculate the ro
	for _, seat := range seats {
		rowNum, seatNum := getSeatNumber(strings.ToLower(seat))
		id := rowNum*8 + seatNum
		if id > best {
			best = id
		}
	}
	fmt.Println("bestSeat", best)
}

func getSeatNumber(seat string) (int, int) {
	return getColumnNumberV2(seat[0:7]), getSeatPositionV2(seat[7:10])
}

func getColumnNumberV2(seat string) int {

	frontRow, backRow := 0, 127

	for i, char := range seat {
		if string(char) == front {
			backRow -= (backRow - frontRow + 1) / 2
			if i == 6 {
				return frontRow
			}
		} else if string(char) == back {
			frontRow += (backRow - frontRow + 1) / 2
			if i == 6 {
				return backRow
			}
		}

	}
	return 0
}

func getSeatPositionV2(seat string) int {
	leftSeat := 0
	rightSeat := 7

	for i, char := range seat {

		if string(char) == left {
			rightSeat -= (rightSeat - leftSeat + 1) / 2
		} else if string(char) == right {
			leftSeat += (rightSeat - leftSeat + 1) / 2
		}

		if i == 2 {
			if string(char) == left {
				return leftSeat
			}
		}

	}

	return rightSeat
}

func readInput() []string {
	f, _ := os.Open("./data.txt")

	scanner := bufio.NewScanner(f)

	var seats []string
	for scanner.Scan() {
		seats = append(seats, scanner.Text())
	}

	return seats
}
