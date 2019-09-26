package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"

	"github.com/vyskocilm/postkid/builder"
)

var (
	verbose bool
	runCurl bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "be verbose")
	flag.BoolVar(&runCurl, "run-curl", false, "execute the command line in curl")
}

func processFiles(files []string) {

    b := builder.New()

	for _, file := range files {
		req, err := builder.ParseFile(file)
		if err != nil {
			log.Fatal(err)
		}
		if !runCurl {
			fmt.Printf(b.Curl(&req))
		} else {
			cmd := b.CurlCmd(&req)
			var out bytes.Buffer
			var sterr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &sterr
			err := cmd.Run()
			if err != nil {
                log.Printf("curl failed, stderr:\n%s", string(sterr.Bytes()))
				log.Fatal(err)
			}
            fmt.Println(string(out.Bytes()))
		}
	}
}

func main() {

	flag.Parse()
	processFiles(flag.Args())

}
