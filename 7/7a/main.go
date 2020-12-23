package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func checkContains(bagStr string, inFile string) int {
	containsCount := 0
	for _, l := range strings.Split(string(inFile), "\n") {
		containsList := strings.Split(l, "bags contain")
		if strings.Contains(containsList[1], bagStr) {
			containsCount += 1 + checkContains(containsList[0], inFile)
		}
	}
	return containsCount
}

func main() {
	inFile, err := ioutil.ReadFile("../input2")
	if err != nil {
		log.Fatal(err)
	}

	bagStr := "shiny gold"
	fmt.Println(checkContains(bagStr, string(inFile)))
}
