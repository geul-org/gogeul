//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what codecDocOf — test fixture: a codec spec document declaring target with raw fields of the given widths
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// codecDocOf builds a spec document whose one DECLARE gives target raw fields
// of the given widths (FIELD Edge Ids 2, 3, …).
func codecDocOf(target core.Word, widths ...core.Word) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, uint32(target))
	in := append(enc(core.EncodeSpecVersion(core.SpecVersion)), enc(id)...)
	in = append(in, 0x1C30)
	in = append(in, enc(core.EncodeDeclare(core.Declare{EdgeID: 1, Target: target}))...)
	for i, w := range widths {
		in = append(in, enc(core.EncodeFieldDecl(core.FieldDecl{EdgeID: uint64(i + 2), Declare: 1, Width: w}))...)
	}
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
