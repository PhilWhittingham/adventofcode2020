package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func checkContains(bagStr string, inFile string) []int {
	containsIndexList := []int{}
	for i, l := range strings.Split(string(inFile), "\n") {
		containsList := strings.Split(l, "bags contain")
		if strings.Contains(containsList[1], bagStr) {
			containsIndexList = append(containsIndexList, i)
			containsIndexList = append(containsIndexList, checkContains(containsList[0], inFile)...)
		}
	}
	return containsIndexList
}

func main() {
	inFile, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	bagStr := "shiny gold"
	counter := make(map[int]int)
	for _, val := range checkContains(bagStr, string(inFile)) {
		counter[val]++
	}
	fmt.Println(len(counter))
}
