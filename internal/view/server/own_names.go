//ff:func feature=server type=helper control=iteration dimension=1
//ff:what ownNames — name → document ID for the documents of one Issuer in document list bytes; broken bytes give none
package server

import "github.com/geul-org/gogeul/pkg/resolve"

// ownNames parses b and keeps the names whose document belongs to space.
func ownNames(b []byte, space string) map[string]string {
	out := map[string]string{}
	names, _ := resolve.ParseDocumentNames(b)
	for n, d := range names {
		if d.Issuer == space {
			out[n] = d.Doc
		}
	}
	return out
}
