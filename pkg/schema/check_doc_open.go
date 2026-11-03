//ff:func feature=schema type=rule control=sequence
//ff:what checkDocOpen — D1: the opening DOC-* record, its length and its narrowest width
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkDocOpen applies D1 to the first record: a DOC-1/2/4WORD Meta whose
// payload is the whole document's word count at the narrowest width.
func checkDocOpen(r rec, words int) []Fact {
	if r.R.Kind != core.KindMeta || r.R.Codec > core.MetaDoc4Word {
		return []Fact{factAt("D1", Error, r, "TAG WORD", "DOC-1WORD/2WORD/4WORD", fmt.Sprintf("kind %d codec 0x%02X", r.R.Kind, r.R.Codec))}
	}
	w := core.WidthWords(uint8(r.R.Codec))
	if len(r.R.Payload) != w {
		return []Fact{factAt("D1", Error, r, "Payload", fmt.Sprintf("%d WORD", w), fmt.Sprintf("%d WORD", len(r.R.Payload)))}
	}
	total, _ := core.DecodeID(r.R.Payload, w)
	if total != uint64(words) {
		return []Fact{factAt("D1", Error, r, "Payload", fmt.Sprintf("문서 전체 %d WORD", words), fmt.Sprintf("%d", total))}
	}
	if w > core.MinWidth(total) {
		return []Fact{factAt("D1", Error, r, "TAG WORD", fmt.Sprintf("DOC-%dWORD", core.MinWidth(total)), fmt.Sprintf("DOC-%dWORD", w))}
	}
	return nil
}
