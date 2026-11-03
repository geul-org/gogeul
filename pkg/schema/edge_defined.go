//ff:func feature=schema type=rule control=selection
//ff:what edgeDefined — whether an Edge codec (code - 0x80) is assigned (§3.14 codes included) or an EDGE-ISSUER code
package schema

import "github.com/geul-org/gogeul/pkg/core"

// edgeDefined reports whether c (Edge Codec = code - 0x80) is an assigned
// direct Edge code (grammar Edge tables; 0x80~0x94 are the TOULMIN record types
// moved from GEUL ISSUER formats, TAG 0x100L~0x114L, migration.md §3.14) or an
// EDGE-ISSUER code (0x13D~0x13F). Reserved 0x7C~0x7F and the empty 0x95~0x13C
// are not.
func edgeDefined(c uint16) bool {
	switch {
	case c <= 0x06, c >= 0x08 && c <= 0x0B, c >= 0x10 && c <= 0x12, c >= 0x18 && c <= 0x1E:
		return true
	case c == 0x20, c == 0x41, c == 0x42, c >= 0x48 && c <= 0x59:
		return true
	case c >= 0x60 && c <= 0x65, c >= 0x68 && c <= 0x6A, c >= 0x70 && c <= 0x76, c == 0x78, c == 0x79:
		return true
	case c >= 0x80 && c <= 0x94:
		return true
	case c >= core.CodecEdgeIssuer2Word && c <= core.CodecEdgeIssuer16Word:
		return true
	default:
		return false
	}
}
