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
			runePresent := true
			for _, f := range strings.Fields(l) {
				if !strings.ContainsRune(f, r) {
					runePresent = false
				}
			}
			if runePresent {
				ansSum++
			}
		}
	}
	fmt.Println(ansSum)
}
