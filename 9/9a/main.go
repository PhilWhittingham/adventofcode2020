package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func isSumInSlice(numSlice []int, sum int) bool {
	for _, i := range numSlice {
		for _, j := range numSlice {
			if i+j == sum {
				return true
			}
		}
	}
	return false
}

func main() {
	inFile, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	preambleVals := []int{}
	for i, l := range strings.Split(string(inFile), "\n") {
		val, _ := strconv.Atoi(l)
		if i < 25 {
			preambleVals = append(preambleVals, val)
		} else {
			if !isSumInSlice(preambleVals, val) {
				fmt.Println(val)
			}
			_, preambleVals = preambleVals[0], preambleVals[1:]
			preambleVals = append(preambleVals, val)
		}

	}
}
