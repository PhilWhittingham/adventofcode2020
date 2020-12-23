package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func valInSlice(list []int, val int) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	inFile, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	acc := 0
	loc := 0
	execLines := []int{}
	fileLines := strings.Split(string(inFile), "\n")
	for !valInSlice(execLines, loc) {
		execLines = append(execLines, loc)
		comVal := strings.Fields(fileLines[loc])
		com := comVal[0]
		val, _ := strconv.Atoi(comVal[1])
		switch com {
		case "nop":
			loc++
		case "acc":
			acc += val
			loc++
		case "jmp":
			loc += val
		}
	}
	fmt.Println(acc)
}
