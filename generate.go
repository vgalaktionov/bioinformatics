package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const implTempl = `
package exercises

func foobar(a []byte, b []byte) int {
	return 1
}

func (ex Exercise) %s(params [][]byte) int {
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

	"github.com/vgalaktionov/bioinformatics/exercises"
)

func Test%s(t *testing.T) {
	expected := "%s"
	result := exercises.Exercise{}.%s([][]byte{%s})

	if result != expected {
		t.Errorf("expected: %%d; got: %%d", expected, result)
	}
}

func Benchmark%s(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.%s([][]byte{%s})
	}
}
`

func Generate() {
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
		datasetInputs = append(datasetInputs, fmt.Sprintf(`[]byte("%s")`, p))
	}
	joined := strings.Join(datasetInputs, ", ")

	// write files to disk
	ioutil.WriteFile(fmt.Sprintf("./exercises/%s.go", ex), []byte(fmt.Sprintf(implTempl, strings.ToUpper(ex))), 0777)

	ioutil.WriteFile(fmt.Sprintf("./exercises/%s_test.go", ex), []byte(fmt.Sprintf(testTempl, exUpper, output, exUpper, joined, exUpper, exUpper, joined)), 0777)
}
