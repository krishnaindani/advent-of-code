package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	data map[string]bagContains
}

type bagContains struct {
	contains map[string]int
}

var colorToFind = "shiny gold"

func main() {
	rawData := readInput()

	rules := parseRules(rawData)
	fmt.Println("rules", rules)

	//b := getNumberOfBagsThatContainsGivenType(rules, colorToFind)
	//fmt.Println("b", b)
}

func readInput() []string {
	f, err := os.Open("./sampleData.txt")
	if err != nil {
		panic(err)
	}

	var rules []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}
	//rules = append(rules, "")
	return rules
}

func parseRules(rules []string) bag {

	b := bag{
		data: make(map[string]bagContains),
	}

	for _, v := range rules {

		//if the bag is empty
		if strings.Contains(v, "bags contain no other bags") {
			topBag := strings.Split(v, "bags contain no other bags")
			b.data[topBag[0]] = bagContains{}
			continue
		}

		//if the bag is not empty
		topBag := strings.Split(v, "bags contain")
		childBags := topBag[1]
		childBags = strings.ReplaceAll(childBags, "bags", "bag")
		bc := bagContains{}

		containsMap := make(map[string]int)

		//it contains multiple bags here
		if strings.Contains(childBags, ",") {
			c := strings.Split(childBags, "bag,")

			for _, v1 := range c {
				c1 := strings.Split(v1, " ")
				color := fmt.Sprintf("%s %s", c1[2], c1[3])
				count, _ := strconv.Atoi(c1[1])
				containsMap[color] = count
			}
			bc.contains = containsMap
		} else {
			//it contains only one type of bag
			c := strings.ReplaceAll(childBags, "bag.", "")
			c1 := strings.Split(c, " ")
			color := fmt.Sprintf("%s %s", c1[2], c1[3])
			count, _ := strconv.Atoi(c1[1])
			containsMap[color] = count
			bc.contains = containsMap
		}
		b.data[topBag[0]] = bc
	}

	return b
}

func getNumberOfBagsThatContainsGivenType(bags bag, bagName string) int {

	var count int

	for _, v := range bags.data {
		_ = v
	}
	return count
}

func isBagPresent() {

}
