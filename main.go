package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

    "github.com/vyskocilm/postkid/builder"

	yaml "gopkg.in/yaml.v2"
)

var (
	verbose bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "be verbose")
}



func processFile(file string) {

	r, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	var req builder.Request
	err = yaml.Unmarshal(b, &req)
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("req=%+v", req)

	fmt.Printf("%s\n", builder.Curl(&req))
}

func processFiles(files []string) {
	for _, file := range files {
		processFile(file)
        req, err := builder.ParseFile(file)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf(builder.Curl(&req))
	}
}

func main() {

	flag.Parse()
	processFiles(flag.Args())

}
