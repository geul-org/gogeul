//ff:func feature=schema type=decoder control=selection
//ff:what readLength — a record's total words and TAG words from its Length code, with canonical and trust flags
package schema

import "github.com/geul-org/gogeul/pkg/core"

// readLength decodes the Length of the TAG WORD at ws[at] (README Length). It
// reports whether the encoding is canonical (F3) and whether the length can be
// trusted at all: an extension must be present and cover its own TAG words.
// The canonical bounds match the encoder: a 1-word extension starts at 16
// words (15 would fit a direct code without the extension word) and a 2-word
// extension at 0x10001 (0x10000 would fit the 1-word extension).
func readLength(ws []core.Word, at int) (total, head int, canonical, ok bool) {
	code := ws[at] & 0xF
	switch {
	case code < 0xE:
		return int(code) + 1, 1, true, true
	case code == 0xE && at+1 < len(ws):
		t := int(ws[at+1])
		return t, 2, t >= 16, t >= 2
	case code == 0xF && at+2 < len(ws):
		t := int(ws[at+1])<<16 | int(ws[at+2])
		return t, 3, t >= 0x10001, t >= 3
	default:
		return 0, 0, false, false
	}
}
