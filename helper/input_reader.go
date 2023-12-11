package helper

import "strings"

func GetInputRows(input string) []string {
	return strings.Split(input, "\n")
}
