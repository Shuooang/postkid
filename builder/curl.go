package builder

import (
	"net/url"
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
