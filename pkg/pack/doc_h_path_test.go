//ff:func feature=pack type=helper control=sequence
//ff:what docHPath — the hashes/ path of a document by its content hash, under the Issuer its DOC-ID names (the GEUL Issuer without one)
package pack

import "github.com/geul-org/gogeul/pkg/core"

// docHPath names doc by its content hash under its own Issuer (its DOC-ID),
// or the GEUL Issuer 0x00000000 for a document without DOC-ID.
func docHPath(doc []byte) string {
	h, _ := contentHash(doc)
	is := "0x00000000"
	if ws, ok := docIssuer(doc); ok {
		is = core.FormatID(ws)
	}
	return hPath(is, string(h[:]))
}
