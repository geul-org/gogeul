//ff:func feature=server type=helper control=sequence
//ff:what idFolder — the folder an ID of the document list lives in: codecs for an 8-digit codec ID, docs for a 16-digit document ID
package server

// idFolder is `codecs` for a codec ID (8 hex digits) and `docs` for a
// document ID (16 hex digits).
func idFolder(id string) string {
	if len(id) == 2+8 {
		return "codecs"
	}
	return "docs"
}
