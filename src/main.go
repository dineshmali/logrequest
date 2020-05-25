package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args
	CaseNumber := input[1]

	fmt.Println(CaseNumber)

	check := contains(input, "sss")
	fmt.Println(check)
}

func contains(array []string, str string) bool {
	for _, a := range array {
		if a == str {
			return true
		}
	}
	return false
}
