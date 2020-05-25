package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args
	CaseNumber := input[1]
	str := "sss"
	check := contains(input, str)
	if check {
		words := []string{CaseNumber, "SHOW", "SUB", "SUM"}
		joined := strings.Join(words, "_")
		wordsnew := []string{"SR", joined}
		joined = strings.Join(wordsnew, "")
		fmt.Printf("ssh user@[Controller0IP] SHOW SUB SUB > %s_C0.txt\n", joined)
		fmt.Printf("ssh user@[Controller1IP] SHOW SUB SUB > %s_C1.txt\n", joined)
	}

}

func contains(array []string, str string) bool {
	for _, a := range array {
		if a == str {
			return true
		}
	}
	return false
}
