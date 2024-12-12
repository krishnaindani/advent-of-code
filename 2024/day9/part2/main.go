package main

import (
	"bufio"
	"fmt"
	"github/krishnaindani/advent-of-code/utils"
	"time"

	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer utils.TimeTrack(time.Now())
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("computeCheckSum:", computeCheckSum(input))
}

func computeCheckSum(nums []int) int {
	block := getBlock(nums)
	return getCheckSumResult(block)
}

func getCheckSumResult(nums []int) int {

	maxFileId := -1
	n := len(nums) - 1

	for _, block := range nums {
		if block > maxFileId {
			maxFileId = block
		}
	}

	for fileID := maxFileId; fileID >= 0; fileID-- {

		fileEndPos := -1
		//find the file position
		for i := n; i >= 0; i-- {
			if nums[i] == fileID {
				fileEndPos = i
				break
			}
		}

		if fileEndPos == -1 {
			continue
		}

		fileStartPos := findFileStartPosition(fileEndPos, nums)
		fileSize := fileEndPos - fileStartPos + 1

		bestFreeSpace := -1

		for i := 0; i < fileStartPos; i++ {
			if nums[i] == -1 {
				freeSize := findFreeSpaceSize(i, nums)
				if freeSize >= fileSize {
					bestFreeSpace = i
					break
				}
			}
		}

		if bestFreeSpace != -1 {
			moveFile(&nums, bestFreeSpace, fileStartPos, fileSize)
		}
	}

	return calculateCheckSum(nums)
}

func calculateCheckSum(nums []int) int {

	result := 0

	for i, num := range nums {
		if num != -1 {
			result += i * num
		}
	}

	return result
}

func moveFile(nums *[]int, freeSizeStartPos, fileStartPos, size int) *[]int {

	id := (*nums)[fileStartPos]

	for i := fileStartPos; i < fileStartPos+size; i++ {
		(*nums)[i] = -1
	}

	for i := freeSizeStartPos; i < freeSizeStartPos+size; i++ {
		(*nums)[i] = id
	}

	return nil
}

func findFreeSpaceSize(position int, nums []int) int {
	size := 0

	for i := position; i < len(nums) && nums[i] == -1; i++ {
		size += 1
	}

	return size
}

func findFileStartPosition(pos int, nums []int) int {

	fileId := nums[pos]
	start := pos

	for start >= 0 && nums[start] == fileId {
		start -= 1
	}

	return start + 1
}

func getBlock(nums []int) []int {

	id := 0
	blocks := make([]int, 0)

	for i, num := range nums {
		if i%2 == 0 {
			blocks = append(blocks, buildBlock(id, num)...)
			id += 1
		} else {
			blocks = append(blocks, buildEmptyBlock(num)...)
		}
	}

	return blocks
}

func buildBlock(num int, frequency int) []int {

	block := make([]int, 0)

	for i := 0; i < frequency; i++ {
		block = append(block, num)
	}

	return block
}

func buildEmptyBlock(frequency int) []int {
	blocks := make([]int, 0)

	for i := 0; i < frequency; i++ {
		blocks = append(blocks, -1)
	}

	return blocks
}

func readInput() ([]int, error) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Couldn't open file: %v\n", err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Couldn't close file: %v\n", err)
		}
	}(file)

	nums := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for _, v := range line {
			i, _ := strconv.Atoi(v)
			nums = append(nums, i)
		}
	}

	return nums, nil
}
