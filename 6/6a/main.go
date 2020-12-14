package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inFile, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	ansSum := 0
	checkRunes := []rune{'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for _, l := range strings.Split(string(inFile), "\n\n") {
		for _, r := range checkRunes {
			if strings.ContainsRune(l, r) {
				ansSum++
			}
		}
	}
	fmt.Println(ansSum)
}
