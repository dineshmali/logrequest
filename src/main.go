package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args
	CaseNumber := input[1]
	CusName := input[2]
	var checktxt bool
	var CMD string
	if "sss" == contains(input, "sss") {
		CMD = " SHOW SUB SUM"
		printer(CaseNumber, CMD, CusName)
		checktxt = true
	}
	if "sssa" == contains(input, "sssa") {
		CMD = " SHOW SUB SUM ALL"
		printer(CaseNumber, CMD, CusName)
		checktxt = true
	}
	if "asss" == contains(input, "asss") {
		CMD = " APP SHOW SUB SUM"
		printer(CaseNumber, CMD, CusName)
		checktxt = true
	}
	if "asssa" == contains(input, "asssa") {
		CMD = " APP SHOW SUB SUM ALL"
		printer(CaseNumber, CMD, CusName)
		checktxt = true
	}

	if "diag" == contains(input, "diag") {
		if "report" == contains(input, "report") {
			diagprinter(CaseNumber, CusName, "report")
		}
		if "lite" == contains(input, "lite") {
			diagprinter(CaseNumber, CusName, "lite")
		} else {
			diagprinter(CaseNumber, CusName, "")
		}
	}
	if checktxt {
		a := "SR" + CaseNumber + " " + CusName
		s := strings.Split(a, " ")
		joined := strings.Join(s, "_")
		fmt.Println("\n Before uploading add text files to single archive:")
		fmt.Printf("\n tar -czf %s_logs.tgz *.txt \n\n", joined)
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

func printer(CaseNumber string, CMD string, cusname string) {
	a := "SR" + CaseNumber + " " + cusname + CMD
	s := strings.Split(a, " ")
	joined := strings.Join(s, "_")
	fmt.Printf("ssh user@[Controller0IP]%s > %s_C0.txt\n", CMD, joined)
	fmt.Printf("ssh user@[Controller1IP]%s > %s_C1.txt\n", CMD, joined)
}

func diagprinter(x string, y string, z string) {
	fmt.Printf("ssh diag@[Controller0IP] %s tgz > SR%s_%s_diag%s_C0.tgz\n", z, x, y, z)
	fmt.Printf("ssh user@[Controller1IP] %s tgz > SR%s_%s_diag%s_C0.tgz\n", z, x, y, z)
}
