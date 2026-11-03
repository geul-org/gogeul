//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what EncodeDocument — a whole document from HEAD and BODY records: DOC framing, the BODY marker when the HEAD has records, END
package core

// EncodeDocument encodes head, the BODY marker, body and END and wraps them in
// the DOC record of the narrowest width (WrapDocument). An empty head omits
// the BODY marker (document.md, 0x1C3L BODY: "HEAD가 비어 있으면 BODY를
// 생략한다"), so the same content has one encoding and one content hash. The
// first record that does not encode is the error.
func EncodeDocument(head, body []Record) ([]byte, error) {
	recs := append([]Record{}, head...)
	if len(head) > 0 {
		recs = append(recs, Record{Kind: KindMeta, Codec: MetaBody})
	}
	in := []Word{}
	for _, r := range append(recs, body...) {
		ws, err := EncodeRecord(r)
		if err != nil {
			return nil, err
		}
		in = append(in, ws...)
	}
	end, _ := EncodeRecord(Record{Kind: KindMeta, Codec: MetaEnd})
	return PutWords(WrapDocument(append(in, end...))), nil
}
