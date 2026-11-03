//ff:func feature=schema type=rule control=sequence
//ff:what noLinguisticContent — whether LANGUAGE-BCP47-CODE bytes name zxx (no linguistic content) as their primary language
package schema

import (
	"bytes"
	"strings"
)

// noLinguisticContent reports whether the tag in b (the payload bytes of a
// LANGUAGE-BCP47-CODE, with odd's final padding byte when odd) has the
// primary language subtag zxx, case ignored.
func noLinguisticContent(b []byte, odd bool) bool {
	if odd && len(b) > 0 {
		b = b[:len(b)-1]
	}
	primary, _, _ := bytes.Cut(b, []byte("-"))
	return strings.EqualFold(string(primary), "zxx")
}
