//ff:type feature=pack type=model
//ff:what pathRef — an archive path mapped to its key space: P (Issuer, document, slot), H (content hash) or named
package pack

// pathRef is a path's place in the key spaces (§3.4). issuer is the canonical
// Issuer hex (a zero-padded spelling read as the same Issuer); for spaceP doc
// and slot are the document ID and the slot (0 versions.geul, v a version),
// for spaceH hash is the 32 raw bytes; path is the canonical spelling.
// spaceNone is a path outside the 공표 주소 forms (or latest.geul, which is
// not stored).
type pathRef struct {
	space  int
	issuer string
	doc    uint64
	slot   uint32
	hash   string
	path   string
}
