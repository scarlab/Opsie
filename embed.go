package embedui

import "embed"

// Embed the entire UI dist folder
//go:embed ui/dist/*
var EmbeddedUI embed.FS
