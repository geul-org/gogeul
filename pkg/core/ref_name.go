//ff:func feature=core type=helper control=sequence
//ff:what refName — the spec name of a Reference C: pointer, REFERENCE-ISSUER or 5Q scheme ("" for slot markers)
package core

// refName names the Reference C: a 5Q scheme by C/5 (schemeNames), or a pointer
// or REFERENCE-ISSUER code (refNames). Slot markers are named by
// SlotMarkerName, not here, and return "".
func refName(c uint16) string {
	if s, ok := RefScheme(c); ok {
		return schemeNames[s]
	}
	return refNames[c]
}
