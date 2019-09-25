package builder

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

// Curl return curl command line
func Curl(req *Request) string {

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
		sb.WriteString(strings.TrimSpace(req.Body)) //FIXME: escape '
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

// Curl return exec.Command with arguments for curl
func CurlCmd(req *Request) *exec.Cmd {

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

    return exec.Command("curl", args...)
}
