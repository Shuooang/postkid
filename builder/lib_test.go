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

func TestMe(t *testing.T) {
	assert := assert.New(t)

	req, err := ParseString(yml)
	assert.NoError(err)
    assert.Equal("curl -XGET -H 'ham: spam' 'http://example.com/?foo=bar' ", Curl(&req))
}
