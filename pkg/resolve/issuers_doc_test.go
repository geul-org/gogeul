//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what issuersDoc — test fixture: an issuers.geul binding names to 1-word Issuers, plus a BIND to an ordinary document and a non-BIND record
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// issuersDoc builds an Issuer list. Each pair is {Issuer low word (high word
// 0), name}.
func issuersDoc(pairs ...[2]any) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	bind := func(id uint64, subj core.Record, name string) []core.Word {
		p := append(enc(core.EncodeLocalRef(id)), enc(subj)...)
		p = append(p, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(name)))...)
		return enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p})
	}
	ref := func(issuer, doc core.Word) core.Record {
		if doc == 0 {
			return core.Record{Kind: core.KindReference, Codec: core.CodecIssuerPointer, Payload: []core.Word{0, issuer}}
		}
		return core.Record{Kind: core.KindReference, Codec: core.CodecDocPointer, Payload: []core.Word{0, issuer, 0, 0, 0, doc, 0, 0}}
	}
	in := []core.Word{0x1C30}
	for i, p := range pairs {
		in = append(in, bind(uint64(i+1), ref(p[0].(core.Word), 0), p[1].(string))...)
	}
	in = append(in, bind(90, ref(1, 5), "a-document")...)
	in = append(in, enc(core.EncodeUintValue(7))...)
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
