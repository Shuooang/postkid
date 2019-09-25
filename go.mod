module github.com/vyskocilm/postkid

go 1.12

require (
	github.com/stretchr/testify v1.4.0
	github.com/vyskocilm/postkid/builder v0.0.0-20190925101945-1964dfbe00b7
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/vyskocilm/postkid/builder => ./builder
