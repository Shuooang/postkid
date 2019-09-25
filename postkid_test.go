package main_test

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMe(t *testing.T) {
	assert := assert.New(t)

	gobuild := exec.Command("go", "build")
	err := gobuild.Run()
	assert.NoError(err)

	postkid := exec.Command("./postkid", "examples/test.yaml")
	var out bytes.Buffer
    postkid.Stdout = &out
    err = postkid.Run()
	assert.NoError(err)

    assert.Equal("curl -XGET -H 'ham: spam' 'http://example.com/path?foo=bar' ", string(out.Bytes()))

}
