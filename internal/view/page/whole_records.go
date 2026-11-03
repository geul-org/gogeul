//ff:func feature=page type=helper control=iteration dimension=1
//ff:what wholeRecords — whether a run of at least 2 words is exactly a sequence of named Value or Reference records
package page

import "github.com/geul-org/gogeul/pkg/core"

// wholeRecords reports whether ws (at least 2 words) splits exactly into
// Value or Reference records of named codecs with nothing left over. Slot
// markers (WORD 0x0000, 0x0010, 0x0020) are taken only in a
// declared field that names them, never out of raw words. A single raw word is
// never read as a record: one word decodes as some record too easily.
func wholeRecords(ws []core.Word) bool {
	if len(ws) < 2 {
		return false
	}
	for len(ws) > 0 {
		r, n, err := core.DecodeRecord(ws)
		if err != nil || (r.Kind != core.KindValue && r.Kind != core.KindReference) || core.CodecName(r.Kind, r.Codec) == "" || r.Kind == core.KindReference && core.IsSlotMarkerCodec(r.Codec) {
			return false
		}
		ws = ws[n:]
	}
	return true
}
