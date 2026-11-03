//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what bindDoc — test helper: a document of SPEC-VERSION and BIND Edges (Edge Id 1, 2, …), the shape of the Issuer list
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// bindDoc wraps binds, numbered 1, 2, … in order, as a document with
// SPEC-VERSION in HEAD.
func bindDoc(binds ...core.Bind) []byte {
	ws, _ := core.EncodeRecord(core.EncodeSpecVersion(core.SpecVersion))
	in := append(ws, core.TargetWord(core.KindMeta, core.MetaBody))
	for i, b := range binds {
		b.EdgeID = uint64(i + 1)
		r, _ := core.EncodeBind(b)
		ws, _ := core.EncodeRecord(r)
		in = append(in, ws...)
	}
	return core.PutWords(core.WrapDocument(append(in, core.TargetWord(core.KindMeta, core.MetaEnd))))
}
