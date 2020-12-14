package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-playground/validator"
)

type passport struct {
	Byr string `json:"byr" validate:"required"`
	Iyr string `json:"iyr" validate:"required"`
	Eyr string `json:"eyr" validate:"required"`
	Hgt string `json:"hgt" validate:"required"`
	Hcl string `json:"hcl" validate:"required"`
	Ecl string `json:"ecl" validate:"required"`
	Pid string `json:"pid" validate:"required"`
	Cid string `json:"cid"`
}

func main() {
	inFile, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	v := validator.New()

	validCount := 0
	for _, l := range strings.Split(string(inFile), "\n\n") {
		fields := strings.Fields(l)

		if len(fields) <= 6 {
			continue
		} else {
			jsonFields := strings.Join(fields, "\",\"")
			jsonFields = strings.Replace(jsonFields, ":", "\":\"", -1)
			jsonFields = "{\"" + jsonFields + "\"}"
			p := passport{}

			err = json.Unmarshal([]byte(jsonFields), &p)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = v.Struct(p)
			if err == nil {
				validCount++
			}
		}
	}
	fmt.Println(validCount)
}
