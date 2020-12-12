package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type rule struct {
	right int
	down  int
}

func main() {
	inFile, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	rules := []rule{
		rule{
			right: 1,
			down:  1,
		},
		rule{
			right: 3,
			down:  1,
		},
		rule{
			right: 5,
			down:  1,
		},
		rule{
			right: 7,
			down:  1,
		},
		rule{
			right: 1,
			down:  2,
		},
	}

	treeMult := 1
	for _, r := range rules {
		tree := '#'
		position := 0
		treeCount := 0
		skip := false
		for _, l := range strings.Split(string(inFile), "\n") {
			if len(l) > 0 {
				if !skip {
					if rune(l[position%len(l)]) == tree {
						treeCount++
					}
					position += r.right

					if r.down == 2 {
						skip = true
					}
				} else {
					skip = false
				}
			}

		}
		treeMult = treeMult * treeCount
	}
}
