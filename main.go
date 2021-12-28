package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func ex() {
	ex := os.Args[2]
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

const implTempl = `
package exercises

func foobar(a string, b string) int {
	return 1
}

func (ex Exercise) %s(params []string) int {
	a := params[0]
	b := params[1]

	result := foobar(a, b)

	return result
}
`

const testTempl = `
package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func Test%s(t *testing.T) {
	expected := "%s"
	result := exercises.Exercise{}.%s([]string{%s})

	assert.ElementsMatch(t, result, expected)
}

func Benchmark%s(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.%s([]string{%s})
	}
}
`

func generate() {
	ex := os.Args[2]
	exUpper := strings.ToUpper(ex)
	log.Printf("generating exercise: %s\n", ex)
	// Request the HTML page.
	res, err := http.Get(fmt.Sprintf("https://rosalind.info/problems/%s", ex))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// query for test dataset
	var dataset string
	doc.Find("#sample-dataset + .codehilite>pre").Each(func(i int, s *goquery.Selection) {
		dataset += s.First().Text()
	})

	// query for return dataset
	var output string
	doc.Find("#sample-output + .codehilite>pre").Each(func(i int, s *goquery.Selection) {
		output += s.First().Text()
	})
	output = strings.Trim(output, " \n")

	// format inputs for generated testcase
	datasetInputs := []string{}
	for _, p := range strings.Fields(dataset) {
		datasetInputs = append(datasetInputs, fmt.Sprintf(`"%s"`, p))
	}
	joined := strings.Join(datasetInputs, ", ")

	// write files to disk
	ioutil.WriteFile(fmt.Sprintf("./exercises/%s.go", ex), []byte(fmt.Sprintf(implTempl, strings.ToUpper(ex))), 0777)

	ioutil.WriteFile(fmt.Sprintf("./exercises/%s_test.go", ex), []byte(fmt.Sprintf(testTempl, exUpper, output, exUpper, joined, exUpper, exUpper, joined)), 0777)
}

func main() {
	log.Println("bioinformatics algoritms CLI")
	switch os.Args[1] {
	case "gen":
		generate()
	case "ex":
		ex()
	default:
		log.Fatalln("invalid command passed")
	}
}
