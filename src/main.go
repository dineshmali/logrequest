package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args
	CaseNumber := input[1]

	for i := 2; i < len(input); i++ {
		//	fmt.Println(input[i])
		switch true {
		case "sss" == contains(input, "sss"):
			CMD := " SHOW SUB SUM"
			printer(CaseNumber, CMD)
		case "sssa" == contains(input, "sssa"):
			CMD := " SHOW SUB SUM ALL"
			printer(CaseNumber, CMD)
		case "asss" == contains(input, "asss"):
			CMD := " APP SHOW SUB SUM"
			printer(CaseNumber, CMD)
		case "asssa" == contains(input, "asssa"):
			CMD := " APP SHOW SUB SUM ALL"
			printer(CaseNumber, CMD)
		}
	}
}

func contains(array []string, str string) string {
	for _, a := range array {
		if a == str {
			return str
		}
	}
	var none string
	return none
}

func printer(CaseNumber string, CMD string) {
	a := "SR" + CaseNumber + CMD
	s := strings.Split(a, " ")
	joined := strings.Join(s, "_")
	fmt.Printf("ssh user@[Controller0IP]%s > %s_C0.txt\n", CMD, joined)
	fmt.Printf("ssh user@[Controller1IP]%s > %s_C1.txt\n", CMD, joined)
}
