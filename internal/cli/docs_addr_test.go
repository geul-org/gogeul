//ff:func feature=cli type=helper control=sequence
//ff:what docsAddr — the 공표 주소 folder of a document key `<Issuer>/<문서>`: `<Issuer>/docs/<문서>`
package cli

import "strings"

// docsAddr turns the document key `<Issuer>/<문서>` the other commands take
// into the 공표 주소 folder geul show reads (`<Issuer>/docs/<문서>`).
func docsAddr(key string) string {
	is, doc, _ := strings.Cut(key, "/")
	return is + "/docs/" + doc
}
