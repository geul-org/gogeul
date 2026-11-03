//ff:func feature=resolve type=helper control=sequence
//ff:what hashPath — `<Issuer>/hashes/<내용 해시>.geul`, a document's address by content hash under its Issuer
package resolve

import "fmt"

// hashPath is the immutable address of a document of issuer (canonical hex)
// by its content hash (§3.13): `<Issuer>/hashes/0x<64 hex>.geul`. There is no
// global hash address; the Issuer must be known.
func hashPath(issuer string, h [32]byte) string {
	return issuer + "/" + hashesDir + "/0x" + fmt.Sprintf("%X", h[:]) + geulExt
}
