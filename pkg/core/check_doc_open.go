//ff:func feature=core type=decoder control=sequence
//ff:what checkDocOpen — validate the opening DOC-* record against the input length and canonical width
package core

import (
	"errors"
	"fmt"
)

// checkDocOpen validates the opening record of a document: it must be a Meta
// DOC-1WORD/2WORD/4WORD whose 1/2/4-word payload is the total word count of the
// input (one document per input, nothing after END) and whose width is the
// narrowest that holds that total.
func checkDocOpen(r Record, inputWords int) error {
	if r.Kind != KindMeta || r.Codec > MetaDoc4Word {
		return errors.New("geul/core: document must open with DOC-1WORD/2WORD/4WORD")
	}
	w := WidthWords(uint8(r.Codec))
	if len(r.Payload) != w {
		return fmt.Errorf("geul/core: DOC payload must be %d words", w)
	}
	total, _ := DecodeID(r.Payload, w)
	if total != uint64(inputWords) {
		return fmt.Errorf("geul/core: DOC length %d does not match input %d words", total, inputWords)
	}
	if w > MinWidth(total) {
		return errors.New("geul/core: non-canonical DOC width (a narrower DOC-* holds the length)")
	}
	return nil
}
