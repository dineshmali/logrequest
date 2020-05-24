package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args

	CaseNumber := input[1]

	fmt.Println(CaseNumber)

	switch checkword {
	case checkword(input, sss):
		println("sss asked")
	}
}

func checkword(input []string, word string) bool {
	for i, n := range input {
		if n == word {
			return true
		}
	}
	return false
}
