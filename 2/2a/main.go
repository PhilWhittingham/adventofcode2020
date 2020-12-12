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
	rLow, _ := strconv.Atoi(r[0])
	rHigh, _ := strconv.Atoi(r[1])

	rCount := 0
	for _, c := range password {
		if c == char {
			rCount++
		}
	}
	if rCount >= rLow && rCount <= rHigh {
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
