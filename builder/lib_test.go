package builder

import (
	"testing"

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

	b := New(WithCurl("cURL.exe"))
    s, err := b.Curl(&req)
    assert.NoError(err)
	assert.Equal(
        "cURL.exe -XGET -H 'ham: spam' 'http://example.com/?foo=bar' ",
        s)
}

func TestGo(t *testing.T) {
	assert := assert.New(t)

    b := New()
    req := Request {
        Method: "GET",
        Query: map[string]string{"foo": "bar"},
    }
    s, err := b.Go(&req)
    assert.NoError(err)
    assert.Equal("", s)
}
