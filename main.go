package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vyskocilm/postkid/builder"
)

var (
	verbose bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "be verbose")
}

func processFiles(files []string) {
	for _, file := range files {
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
