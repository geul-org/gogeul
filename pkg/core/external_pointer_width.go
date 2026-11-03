//ff:func feature=core type=helper control=selection
//ff:what ExternalPointerWidth — the IMPORT ID width of an EXTERNAL-POINTER scheme (6, 7, 8 → 1, 2, 4)
package core

// ExternalPointerWidth returns the IMPORT ID width of the Reference scheme (C/5)
// of EXTERNAL-POINTER-1/2/4WORD: scheme 6 is 1, 7 is 2 and 8 is 4 words. Any
// other scheme returns 0.
func ExternalPointerWidth(scheme uint16) int {
	switch scheme {
	case SchemeExternalPointer1Word:
		return 1
	case SchemeExternalPointer2Word:
		return 2
	case SchemeExternalPointer4Word:
		return 4
	default:
		return 0
	}
}
