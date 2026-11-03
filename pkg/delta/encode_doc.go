//ff:func feature=delta type=encoder control=iteration dimension=1
//ff:what encodeDoc — a document from HEAD and BODY records, wrapped in its DOC-* record
package delta

import "github.com/geul-org/gogeul/pkg/core"

// encodeDoc writes HEAD, BODY marker, BODY and END, wrapped in the narrowest
// DOC-* record (core.WrapDocument).
func encodeDoc(head, body []core.Record) []byte {
	var in []core.Word
	for _, r := range head {
		ws, _ := core.EncodeRecord(r)
		in = append(in, ws...)
	}
	in = append(in, core.TargetWord(core.KindMeta, core.MetaBody))
	for _, r := range body {
		ws, _ := core.EncodeRecord(r)
		in = append(in, ws...)
	}
	return core.PutWords(core.WrapDocument(append(in, core.TargetWord(core.KindMeta, core.MetaEnd))))
}
