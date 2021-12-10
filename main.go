package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/vgalaktionov/bioinformatics/exercises"
)

func main() {
	log.Println("bioinformatics algoritms CLI")
	if os.Args[1] == "generate" {
		Generate()
		return
	}
	ex := os.Args[1]
	log.Printf("running exercise: %s\n", ex)

	ep := regexp.MustCompile(`ba(\d+)[a-z]`)

	mss := ep.FindAllStringSubmatch(ex, 1)
	if mss == nil {
		log.Fatalln("invalid exercise")
	}

	ch, _ := strconv.Atoi(mss[0][1])
	log.Printf("running chapter: %d", ch)

	v := reflect.ValueOf(exercises.Exercise{})
	m := v.MethodByName(strings.ToUpper(ex))
	if m.Kind() != reflect.Func {
		log.Fatalln("exercise not implemented")
	}

	params := make([]reflect.Value, 1)
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to parse input: %s", err)
	}
	inputStr := string(input)
	log.Printf("running with input: %s", inputStr)
	params[0] = reflect.ValueOf(strings.Fields(inputStr))

	result := m.Call(params)
	log.Printf("%s", result)
}
