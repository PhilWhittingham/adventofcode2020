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

func executes(comList []string) (bool, int) {
	acc := 0
	loc := 0
	execLines := []int{}
	for !valInSlice(execLines, loc) {
		if loc < 0 {
			return false, acc
		}
		execLines = append(execLines, loc)
		comVal := strings.Fields(comList[loc])
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
		if loc >= len(comList) {
			return true, acc
		}
	}
	return false, acc
}

func main() {
	inFile, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	fileLines := strings.Split(string(inFile), "\n")
	for i, l := range fileLines {
		comVal := strings.Fields(l)
		if comVal[0] == "nop" {
			comVal[0] = "jmp"
		} else if comVal[0] == "jmp" {
			comVal[0] = "nop"
		}

		newList := []string{}
		for j, b := range fileLines {
			if i == j {
				newList = append(newList, strings.Join(comVal, " "))
			} else {
				newList = append(newList, b)
			}
		}

		if exec, val := executes(newList); exec {
			fmt.Println(val)
		}

	}
}
