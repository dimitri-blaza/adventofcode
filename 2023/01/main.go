package main

import (
	"adventofcode/helper"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type indexedString struct {
	index int
	str   string
}

func main() {
	inputRows := helper.GetInputRows(input)

	var result int64
	for _, row := range inputRows {
		result += getRowDigitV2(row)
	}

	fmt.Println(result)
}

func getRowDigit(row string) int64 {
	var digits []string
	for _, s := range strings.Split(row, "") {
		if _, err := strconv.ParseInt(s, 10, 64); err == nil {
			digits = append(digits, s)
		}
	}

	return helper.Rerr(strconv.ParseInt(digits[0]+digits[len(digits)-1], 10, 64))
}

func getRowDigitV2(row string) int64 {
	numbers := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	indexes := map[int]string{}
	for str, nb := range numbers {
		indexes[strings.Index(row, str)] = nb
		indexes[strings.LastIndex(row, str)] = nb
	}

	for i, s := range strings.Split(row, "") {
		if _, err := strconv.ParseInt(s, 10, 64); err == nil {
			indexes[i] = s
		}
	}

	var first, last indexedString
	for i, str := range indexes {
		switch {
		case i == -1:
			continue
		case first.str == "":
			first = indexedString{index: i, str: str}
			last = indexedString{index: i, str: str}
		default:
			if first.index > i {
				first = indexedString{index: i, str: str}
			}
			if last.index < i {
				last = indexedString{index: i, str: str}
			}
		}
	}

	return helper.Rerr(strconv.ParseInt(first.str+last.str, 10, 64))
}
