package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type sumPair struct {
	valA int
	valB int
}

func (s *sumPair) sum() int {
	return s.valA + s.valB
}

func main() {
	inFile, err := os.Open("../input")

	numList := []int{}

	if err != nil {
		log.Fatal(err)
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		s := scanner.Text()
		val, _ := strconv.Atoi(s)

		numList = append(numList, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, i := range numList {
		for _, j := range numList {
			for _, k := range numList {
				if i+j+k == 2020 {
					fmt.Println(i, j, k)
					fmt.Println(i * j * k)
				}
			}
		}
	}
}
