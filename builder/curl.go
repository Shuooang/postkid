package builder

import (
	"fmt"
	"io"
	"net/url"
	"os/exec"
	"strings"
)

// Curl return curl command line
func (b *Builder) Curl(req *Request) error {

	// curl
	io.WriteString(b.w, b.curl)
	// ... method
	io.WriteString(b.w, " -X")
	io.WriteString(b.w, req.Method)

	// ... headers
	for name, value := range req.Header {
		io.WriteString(b.w, " -H '")
		io.WriteString(b.w, name)
		io.WriteString(b.w, ": ")
		io.WriteString(b.w, value)
		io.WriteString(b.w, "'")
	}

	if len(req.Body) > 0 && (req.Method == "POST" || req.Method == "PUT") {
		io.WriteString(b.w, " --data '")
		io.WriteString(b.w, strings.TrimSpace(req.Body)) //FIXME: escape '
		io.WriteString(b.w, "'")
	}

	// ... path
	io.WriteString(b.w, " '")
	io.WriteString(b.w, req.Host)
	io.WriteString(b.w, "/")
	io.WriteString(b.w, url.PathEscape(req.Path))

	// ... query
	query := req.QueryString()
	if query != "" {
		io.WriteString(b.w, "?")
		io.WriteString(b.w, query)
	}
	io.WriteString(b.w, "' ")

	return nil
}

// Curl return exec.Command with arguments for curl
func (b *Builder) CurlCmd(req *Request) (*exec.Cmd, error) {

	args := make([]string, 0)

	// ... method
	args = append(args, fmt.Sprintf("-X%s", req.Method))

	// ... headers
	for name, value := range req.Header {
		args = append(args, "-H")
		args = append(args, fmt.Sprintf("%s: %s", name, value))
	}

	if len(req.Body) > 0 && (req.Method == "POST" || req.Method == "PUT") {
		args = append(args, "--data")
		args = append(args, req.Body)
	}

	// ... path
	path := fmt.Sprintf("%s/%s", req.Host, url.PathEscape(req.Path))

	// ... query
	query := req.QueryString()
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	args = append(args, path)

	return exec.Command(b.curl, args...), nil
}
