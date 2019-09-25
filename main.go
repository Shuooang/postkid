package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var (
	verbose bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "be verbose")
}

type Request struct {
	Method string            `yaml:"method"`
	Host   string            `yaml:"host"`
	Path   string            `yaml:"path"`
	Query  map[string]string `yaml:"query,omitempty"`
	Header map[string]string `yaml:"header,omitempty"`
	Body   string            `yaml:"body,omitempty"`
}

// QueryString return properly escaped query as a single string
// return empty string if there were no query parameters specified
func (req *Request) QueryString() string {
	var sbq strings.Builder
	for name, value := range req.Query {
		sbq.WriteString(url.QueryEscape(name))
		sbq.WriteString("=")
		sbq.WriteString(url.QueryEscape(value))
		sbq.WriteString("&")
	}
	if sbq.Len() == 0 {
		return ""
	}
	s := sbq.String()
	s = s[0 : len(s)-1]
	return s
}

// RenderCurl return curl command line
func RenderCurl(req *Request) string {

	var sb strings.Builder

	// curl
	sb.WriteString("curl")
	// ... method
	sb.WriteString(" -X")
	sb.WriteString(req.Method)

	// ... headers
	for name, value := range req.Header {
		sb.WriteString(" -H '")
		sb.WriteString(name)
		sb.WriteString(": ")
		sb.WriteString(value)
		sb.WriteString("'")
	}

    if len(req.Body) > 0 && (req.Method == "POST" || req.Method == "PUT") {
        sb.WriteString(" --data '")
        sb.WriteString(strings.TrimSpace(req.Body))            //FIXME: escape '
        sb.WriteString("'")
    }

	// ... path
	sb.WriteString(" '")
	sb.WriteString(req.Host)
	sb.WriteString("/")
	sb.WriteString(url.PathEscape(req.Path))

	// ... query
	query := req.QueryString()
	if query != "" {
		sb.WriteString("?")
		sb.WriteString(query)
	}
	sb.WriteString("' ")

	return sb.String()
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

	var req Request
	err = yaml.Unmarshal(b, &req)
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("req=%+v", req)

	fmt.Printf("%s\n", RenderCurl(&req))
}

func processFiles(files []string) {
	for _, file := range files {
		processFile(file)
	}
}

func main() {

	flag.Parse()
	processFiles(flag.Args())

}
