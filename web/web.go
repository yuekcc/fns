package web

import "embed"

var (
	//go:embed index.html dist
	Assets embed.FS
)
