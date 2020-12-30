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

	r = "right"
	l = "left"
)

func main() {
	seats := readInput()

	var best int
	//first calculate the ro
	for _, seat := range seats {
		rowNum, seatNum := getSeatNumber(seat)
		id := (rowNum*8 + seatNum)
		if id > best {
			best = id
		}
	}
	fmt.Println("bestSeat", best)
}

func getSeatNumber(seat string) (int, int) {
	s := strings.Split(strings.ToLower(seat), "")
	rowNumbers := s[0 : len(s)-3]
	seatNumbers := s[len(s)-3:]

	return getColumnNumber(rowNumbers), getSeatPosition(seatNumbers)
}

func getColumnNumber(seat []string) int {

	min := 0
	max := 127

	for i, rune := range seat {

		median := (min + max) / 2

		if string(rune) == front {
			max = median - 1
		} else if string(rune) == back {
			min = median + 1
		}

		if i == 6 {
			if string(rune) == front {
				return min
			}
		}

	}

	return max

}

func getSeatPosition(seat []string) int {

	leftSeat := 0
	rightSeat := 7

	for i, rune := range seat {

		if string(rune) == left {
			rightSeat -= (rightSeat - leftSeat + 1) / 2
		} else if string(rune) == right {
			leftSeat += (rightSeat - leftSeat + 1) / 2
		}

		if i == 2 {
			if string(rune) == left {
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
