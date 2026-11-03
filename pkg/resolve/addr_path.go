//ff:func feature=resolve type=helper control=selection
//ff:what Addr.Path — the address relative to `.well-known/geul/` in canonical form (latest.geul, versions.geul, a version or a hash)
package resolve

import "encoding/hex"

// Path writes a back as a path relative to `.well-known/geul/`, the Issuer in
// its canonical width: ParseAddr(a.Path()) gives a again (Base aside).
func (a Addr) Path() string {
	switch {
	case a.Kind == AddrHash:
		return a.folder() + "/0x" + hex.EncodeToString(a.Hash[:]) + geulExt
	case a.Versions:
		return a.folder() + "/" + listFile
	case a.Version == Latest:
		return a.folder() + "/" + latestFile
	default:
		return a.folder() + "/" + FormatVersion(a.Version) + geulExt
	}
}
