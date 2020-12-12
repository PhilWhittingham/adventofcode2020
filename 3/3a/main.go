package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	inFile, err := os.Open("../input")

	if err != nil {
		log.Fatal(err)
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	tree := '#'
	position := 0
	treeCount := 0
	for scanner.Scan() {
		s := scanner.Text()
		if rune(s[position%len(s)]) == tree {
			treeCount++
		}
		position += 3
	}

	println(treeCount)
}
