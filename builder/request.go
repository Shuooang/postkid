package builder

import (
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Request is representation of a request as stored in yaml file
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

func ParseFile(path string) (Request, error) {

	var req Request
	r, err := os.Open(path)
	if err != nil {
		return req, err
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return req, err
	}

	err = yaml.Unmarshal(b, &req)
	if err != nil {
		return req, err
	}
	return req, nil
}

func ParseString(yml string) (Request, error) {
	var req Request
	r := strings.NewReader(yml)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return req, err
	}

	err = yaml.Unmarshal(b, &req)
	if err != nil {
		return req, err
	}
	return req, nil
}
