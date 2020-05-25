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
		CMD := " SHOW SUB SUM"
		printer(CaseNumber, CMD)
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

func printer(CaseNumber string, CMD string) {
	a := "SR" + CaseNumber + CMD
	s := strings.Split(a, " ")
	joined := strings.Join(s, "_")
	fmt.Printf("ssh user@[Controller0IP]%s > %s_C0.txt\n", CMD, joined)
	fmt.Printf("ssh user@[Controller1IP]%s > %s_C1.txt\n", CMD, joined)
}
