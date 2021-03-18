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
	allSeats := make(map[int]bool)

	//first calculate the ro
	for _, seat := range seats {
		rowNum, seatNum := getSeatNumber(seat)
		id := (rowNum*8 + seatNum)
		allSeats[id] = true
	}
	var mySeat int
	fmt.Println("allSeats", allSeats)
	for seatID := range allSeats {
		if allSeats[seatID] && allSeats[seatID+2] && !allSeats[seatID+1] {
			mySeat = seatID + 1
			break
		}
	}

	fmt.Println("mySeat", mySeat)
}

func getSeatNumber(seat string) (int, int) {
	s := strings.Split(strings.ToLower(seat), "")
	rowNumbers := s[0 : len(s)-3]
	seatNumbers := s[7:]

	return getColumnNumber(rowNumbers), getSeatPosition(seatNumbers)
}

func getColumnNumber(seat []string) int {

	columnLeft := 0
	columnRight := 127

	for i, rune := range seat[0:6] {

		if string(rune) == front {
			columnRight -= (columnRight - columnLeft + 1) / 2
		} else if string(rune) == back {
			columnLeft += (columnRight - columnLeft + 1) / 2
		}

		if i == 6 {
			if string(rune) == front {
				return columnLeft
			}
		}

	}

	return columnRight

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
