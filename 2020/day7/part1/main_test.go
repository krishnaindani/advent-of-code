package main

import (
	"fmt"
	"testing"
)

func Test_parseRules(t *testing.T) {

	s := "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags."
	parseRules([]string{s})

	s1 := "bright white bags contain 1 shiny gold bag."
	fmt.Println(parseRules([]string{s1}))
}
