package builder

import (
	"io"
	"os"
)

// Builder contains all the settings and everything
// necessary to construct the output
type Builder struct {
	w    io.Writer // writer to pass output to
	curl string    // path to curl program
}

type Option func(*Builder)

// New returns an instance of a new builder
func New(options ...Option) *Builder {
	b := &Builder{
		w:    os.Stdout,
		curl: "curl",
	}

	for _, o := range options {
		o(b)
	}

	return b
}

// WithOutputWriter sets the writer tool will print an output into
func WithOutputWriter(w io.Writer) Option {
	return func(b *Builder) {
		b.w = w
	}
}

// WithCurl sets the path to curl
func WithCurl(curl string) Option {
	return func(b *Builder) {
		b.curl = curl
	}
}
