package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	assert := assert.New(t)
    b := New()
    assert.Equal("curl", b.curl)

    b2 := New(WithCurl("cURL"))
    assert.Equal("cURL", b2.curl)
}
