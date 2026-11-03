//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what codecDocAs — test fixture: a codec document whose DOC-ID names codec id but whose one DECLARE targets target (a codec of another Kind at that ID)
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// codecDocAs builds a codec document at codec ID id (DOC-ID CODEC-POINTER of
// the GEUL Issuer) declaring target with raw fields of the given widths.
func codecDocAs(id, target core.Word, widths ...core.Word) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	did, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, uint32(id))
	in := append(enc(core.EncodeSpecVersion(core.SpecVersion)), enc(did)...)
	in = append(in, core.TargetWord(core.KindMeta, core.MetaBody))
	in = append(in, enc(core.EncodeDeclare(core.Declare{EdgeID: 1, Target: target}))...)
	for i, w := range widths {
		in = append(in, enc(core.EncodeFieldDecl(core.FieldDecl{EdgeID: uint64(i + 2), Declare: 1, Width: w}))...)
	}
	return core.PutWords(core.WrapDocument(append(in, core.TargetWord(core.KindMeta, core.MetaEnd))))
}
