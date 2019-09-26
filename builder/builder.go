package builder

// Builder contains all the settings and everything
// necessary to construct the output
type Builder struct {
	curl string // path to curl program
}

type Option func(*Builder)

// New returns an instance of a new builder
func New(options ...Option) *Builder {
	b := &Builder{
		curl: "curl",
	}

	for _, o := range options {
		o(b)
	}

	return b
}

// WithCurl sets the path to curl
func WithCurl(curl string) Option {
	return func(b *Builder) {
		b.curl = curl
	}
}
