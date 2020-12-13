package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {

	inFile, err := os.Open("../input")

	if err != nil {
		log.Fatal(err)
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	idList := []int{}
	for scanner.Scan() {
		s := scanner.Text()

		row := strings.Replace(s[0:7], "B", "1", -1)
		row = strings.Replace(row, "F", "0", -1)
		column := strings.Replace(s[7:len(s)], "R", "1", -1)
		column = strings.Replace(column, "L", "0", -1)

		rowInt, _ := strconv.ParseInt(row, 2, 64)
		columnInt, _ := strconv.ParseInt(column, 2, 64)
		
	    idList = append(idList, int((rowInt * 8) + columnInt))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(idList)
	for i, val := range idList {
		if idList[i+1] != val + 1 {
			fmt.Println(idList[i+1] - 1)
			break
		}
	}
}