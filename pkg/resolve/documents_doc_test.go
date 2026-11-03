//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what documentsDoc — test fixture: a documents.geul binding names to documents of Issuer 0x0000
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// documentsDoc builds a documents.geul whose BINDs name document doc (last
// word) of Issuer 0x0000.
func documentsDoc(names map[string]core.Word) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	in := []core.Word{0x1C30}
	id := uint64(1)
	for _, n := range sortedKeys(names) {
		subj := core.Record{Kind: core.KindReference, Codec: core.CodecDocPointer, Payload: []core.Word{0, 0, 0xffff, 0, 0, names[n], 0, 0}}
		p := append(enc(core.EncodeLocalRef(id)), enc(subj)...)
		p = append(p, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(n)))...)
		in = append(in, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p})...)
		id++
	}
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
