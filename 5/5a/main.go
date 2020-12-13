package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {

	inFile, err := os.Open("../input")

	if err != nil {
		log.Fatal(err)
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	var maxID int64 = 0
	for scanner.Scan() {
		s := scanner.Text()

		row := strings.Replace(s[0:7], "B", "1", -1)
		row = strings.Replace(row, "F", "0", -1)
		column := strings.Replace(s[7:len(s)], "R", "1", -1)
		column = strings.Replace(column, "L", "0", -1)

		rowInt, _ := strconv.ParseInt(row, 2, 64)
		columnInt, _ := strconv.ParseInt(column, 2, 64)
		
		if ((rowInt * 8) + columnInt) > maxID {
			maxID = (rowInt * 8) + columnInt
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(maxID)
}