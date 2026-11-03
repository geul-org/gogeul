//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what edgeDoc — test fixture: document 0xffff00000000<id> of Issuer 0x0000 whose BODY is BIND #n = text, numbered from 1
package resolve

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// edgeDoc builds a diffable document: DOC-ID, then BIND(#i, LOCAL #i,
// texts[i-1]) for each text, so a longer version only adds IDs from its
// list's NEXT-LOCAL-ID on.
func edgeDoc(id core.Word, texts ...string) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	doc, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, id})
	in := append(enc(doc), 0x1C30)
	for i, s := range texts {
		p := slices.Concat(enc(core.EncodeLocalRef(uint64(i+1))), enc(core.EncodeLocalRef(uint64(i+1))), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(s))))
		in = append(in, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p})...)
	}
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
