package assets

import "embed"

//go:embed lua-template/*
var TemplateFS embed.FS
