//ff:func feature=resolve type=helper control=sequence
//ff:what codecTargetOf — the target word of a GEUL direct codec ID (8 hex digits): high word 0, a TAG with L=0 that is not a slot marker
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// codecTargetOf reads a codec ID in canonical hex (8 digits, A-F upper case as stored) and returns its
// target word when it is a GEUL direct codec ID (codecDocID): high word 0, then
// a TAG of prefix 000 with LLLL 0000 that is not a slot marker (0x0000,
// 0x0010, 0x0020 have no codec document).
func codecTargetOf(idHex string) (core.Word, bool) {
	b, ok := storedHexID(idHex, 2)
	if !ok {
		return 0, false
	}
	ws, _ := core.ReadWords(b)
	t := ws[1]
	if ws[0] != 0 || t&0xE00F != 0 || core.IsSlotMarkerWord(t) {
		return 0, false
	}
	return t, true
}
