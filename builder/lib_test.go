package builder

import (
	"testing"
    "strings"

	"github.com/stretchr/testify/assert"
)

const yml = `method: GET
host: http://example.com
part: /path
query:
  foo: bar
header:
  ham: spam
`

func TestCurl(t *testing.T) {
	assert := assert.New(t)

	req, err := ParseString(yml)
	assert.NoError(err)

	var sb strings.Builder

	b := New(
		WithCurl("cURL.exe"),
		WithOutputWriter(&sb),
	)
	err = b.Curl(&req)
	assert.NoError(err)
	assert.Equal(
		"cURL.exe -XGET -H 'ham: spam' 'http://example.com/?foo=bar' ",
		sb.String())
}

func TestGo(t *testing.T) {
	assert := assert.New(t)

	var sb strings.Builder

	b := New(
		WithCurl("cURL.exe"),
		WithOutputWriter(&sb),
	)
	req := Request{
		Method: "GET",
		Query:  map[string]string{"foo": "bar"},
	}
    err := b.Go(&req)
	assert.NoError(err)
	assert.NotEqual("", sb.String())
}
