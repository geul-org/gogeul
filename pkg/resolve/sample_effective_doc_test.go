//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what sampleEffectiveDoc — test fixture: a named document with EFFECTIVE-FROM/UNTIL Meta as GREGORIAN-DATE day counts
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// sampleEffectiveDoc builds document 0xffff00000000<id> of Issuer 0x0000 whose
// BODY is text. Each meta is {Meta codec, day count from 1970-01-01}; the day
// count is written as a GREGORIAN-DATE Value inside the Meta.
func sampleEffectiveDoc(id core.Word, text string, meta ...[2]uint32) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	in := enc(func() core.Record {
		r, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, id})
		return r
	}())
	for _, m := range meta {
		v := enc(core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{core.Word(m[1] >> 16), core.Word(m[1])}})
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: uint16(m[0]), Payload: v})...)
	}
	in = append(append(in, 0x1C30), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(text)))...)
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
