package static

import "embed"

// files holds the third-party graph scripts and their licenses that the
// viewer page loads from `_static/` (Write).
//
//go:embed vendor
var files embed.FS
