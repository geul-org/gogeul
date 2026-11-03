//ff:func feature=schema type=rule control=sequence
//ff:what metaDefined — whether a Meta codec is assigned (C 0x00~0x29, no gaps)
package schema

import "github.com/geul-org/gogeul/pkg/core"

// metaDefined reports whether c is an assigned Meta code. The 42 Meta codes
// run from 0x00 (DOC-1WORD) to 0x29 (VERDICT) without gaps; 0x2A~0x3F are
// empty and reserved for validation.
func metaDefined(c uint16) bool {
	return c <= core.MetaVerdict
}
