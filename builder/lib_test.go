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
	assert.Equal("cURL.exe -XGET -H 'ham: spam' 'http://example.com/?foo=bar' ", b.Curl(&req))
}
