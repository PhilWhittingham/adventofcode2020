package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
)

type passport struct {
	Byr string `json:"byr" validate:"required,byr"`
	Iyr string `json:"iyr" validate:"required,iyr"`
	Eyr string `json:"eyr" validate:"required,eyr"`
	Hgt string `json:"hgt" validate:"required,hgt"`
	Hcl string `json:"hcl" validate:"required,hexcolor"`
	Ecl string `json:"ecl" validate:"required,oneof=amb blu brn gry grn hzl oth"`
	Pid string `json:"pid" validate:"required,pid"`
	Cid string `json:"cid"`
}

func main() {
	inFile, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	v := validator.New()
	_ = v.RegisterValidation("byr", func(fl validator.FieldLevel) bool {
		i, _ := strconv.Atoi(fl.Field().String())
		return (i >= 1920) && (i <= 2002)
	})
	_ = v.RegisterValidation("iyr", func(fl validator.FieldLevel) bool {
		i, _ := strconv.Atoi(fl.Field().String())
		return (i >= 2010) && (i <= 2020)
	})
	_ = v.RegisterValidation("eyr", func(fl validator.FieldLevel) bool {
		i, _ := strconv.Atoi(fl.Field().String())
		return (i >= 2020) && (i <= 2030)
	})
	_ = v.RegisterValidation("hgt", func(fl validator.FieldLevel) bool {
		metric := fl.Field().String()[len(fl.Field().String())-2:]
		value, _ := strconv.Atoi(fl.Field().String()[:len(fl.Field().String())-2])
		if metric == "in" {
			return (value >= 59 && value <= 76)
		} else if metric == "cm" {
			return (value >= 150 && value <= 193)
		}
		return false
	})
	_ = v.RegisterValidation("pid", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) == 9
	})

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
