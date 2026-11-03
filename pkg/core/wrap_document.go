//ff:func feature=core type=encoder control=sequence
//ff:what WrapDocument — prefix HEAD+BODY+END words with the narrowest DOC-* record carrying the total length
package core

// WrapDocument opens inner (the HEAD Meta records, BODY, the body records and
// END) with a DOC-1WORD/2WORD/4WORD record whose payload is the document's total
// word count from the DOC record through END (GEUL t002 OVERVIEW 구조). The
// narrowest width that holds the total is used, as the canonical form requires;
// the DOC record's own 2/3/5 words count toward the total. A DOC record is at
// most 5 words, so encoding it cannot fail.
func WrapDocument(inner []Word) []Word {
	w := 4
	if uint64(len(inner))+2 <= 0xFFFF {
		w = 1
	} else if uint64(len(inner))+3 <= 0xFFFFFFFF {
		w = 2
	}
	code, _ := WidthCode(w)
	payload, _ := EncodeID(uint64(len(inner)+1+w), w)
	head, _ := EncodeRecord(Record{Kind: KindMeta, Codec: MetaDoc1Word + uint16(code), Payload: payload})
	return append(head, inner...)
}
