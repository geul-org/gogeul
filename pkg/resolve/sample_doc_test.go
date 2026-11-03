//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what sampleDoc — test fixture: a named document under Issuer 0x0000 with IMPORT-DOC-1WORD records and one text Value
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// sampleDoc builds document 0xffff00000000<id> of Issuer 0x0000 whose BODY is
// text. Each import is {IMPORT ID, imported document id, version ID}.
func sampleDoc(id core.Word, text string, imports ...[3]uint32) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	in := enc(func() core.Record {
		r, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, id})
		return r
	}())
	for _, m := range imports {
		ptr, _ := core.EncodeDocPointer([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, core.Word(m[1])}, []core.Word{core.Word(m[2] >> 16), core.Word(m[2])})
		p := append(enc(ptr), core.EncodeUintValue(uint64(m[0])).Payload...)
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: p})...)
	}
	in = append(append(in, 0x1C30), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(text)))...)
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
