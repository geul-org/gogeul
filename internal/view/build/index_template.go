package build

import "html/template"

// indexTemplate renders a directory listing of the tree mode.
var indexTemplate = template.Must(template.New("index").Parse(`<!doctype html>
<html lang="ko"><head><meta charset="utf-8"><title>{{.Path}} · geul</title>
<style>body{font:14px/1.6 system-ui,sans-serif;margin:24px;color:#1f2328;background:#fff}a{color:#0969da;text-decoration:none}a:hover{text-decoration:underline}li{font-family:ui-monospace,monospace}@media (prefers-color-scheme:dark){body{color:#e6edf3;background:#0d1117}a{color:#58a6ff}}</style></head>
<body><h1>{{.Path}}</h1><ul>{{if .Up}}<li><a href="{{.Up}}">../</a></li>{{end}}{{range .Entries}}<li><a href="{{.URL}}"{{if .ID}} title="{{.ID}}"{{end}}>{{.Name}}</a></li>{{else}}<li>(비어 있음)</li>{{end}}</ul></body></html>
`))
