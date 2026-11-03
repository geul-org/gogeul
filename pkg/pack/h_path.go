//ff:func feature=pack type=helper control=sequence
//ff:what hPath — the address of an H key under an Issuer: `<I>/hashes/<내용 해시>.geul`
package pack

import "encoding/hex"

// hPath is the path of a 32-byte content hash key of issuer (canonical hex).
func hPath(issuer, h string) string {
	return issuer + "/" + hashesDir + "/0x" + hex.EncodeToString([]byte(h)) + ".geul"
}
