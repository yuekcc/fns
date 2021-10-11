package web

import "embed"

var (
	//go:embed dist/*
	Assets embed.FS
)
