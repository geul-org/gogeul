//ff:func feature=resolve type=encoder control=sequence
//ff:what otherDoc — test fixture: document 5 of Issuer 0x0001 with one long, compressible text
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// otherDoc builds the document of a second Issuer, for sets per Issuer.
func otherDoc(text string) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocID([]core.Word{0, 2}, []core.Word{0, 0, 0, 5})
	in := append(enc(id), 0x1C30)
	in = append(in, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(text)))...)
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
