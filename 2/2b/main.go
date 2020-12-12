package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkPassword(rule string, char rune, password string) bool {
	r := strings.Split(rule, "-")
	rInd1, _ := strconv.Atoi(r[0])
	rInd2, _ := strconv.Atoi(r[1])

	rInd1--
	rInd2--
	rCount := 0

	rIndList := []int{}
	if rInd1 < len(password) {
		rIndList = append(rIndList, rInd1)
	}
	if rInd2 < len(password) {
		rIndList = append(rIndList, rInd2)
	}

	for _, rInd := range rIndList {
		if rune(password[rInd]) == char {
			rCount++
		}
	}

	if rCount == 1 {
		return true
	}
	return false
}

func main() {

	inFile, err := os.Open("../input")

	if err != nil {
		log.Fatal(err)
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	passCount := 0
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		if checkPassword(s[0], rune(s[1][0]), s[2]) {
			passCount++
		}
	}
	fmt.Println(passCount)
}
