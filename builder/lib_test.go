package builder

import (
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require := require.New(t)

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

	cwd, err := os.Getwd()
	require.NoError(err)
	err = os.Chdir("testdata")
	require.NoError(err)

	fset := token.NewFileSet() // positions are relative to fset
    f, err := parser.ParseFile(fset, "", sb.String(), parser.AllErrors)
	assert.NoError(err)

    // XXX: this is included only to know ast test works
    if testing.Verbose() {
        t.Logf(sb.String())
        // Print the imports from the file's AST.
        for _, s := range f.Imports {
            t.Logf(s.Path.Value)
        }
    }

	err = os.Chdir(cwd)
	require.NoError(err)
}
