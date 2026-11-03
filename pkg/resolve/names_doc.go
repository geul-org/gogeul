//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what namesDoc — an Issuer's document list with one more BIND: HEAD kept (SPEC-VERSION when new), BIND numbered after the highest Edge Id
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// namesDoc returns the name index head+body with bind appended under the
// next Edge Id. A new index (head nil) gets SPEC-VERSION.
func namesDoc(head, body []core.Record, b core.Bind) ([]byte, error) {
	if head == nil {
		head = []core.Record{core.EncodeSpecVersion(core.SpecVersion)}
	}
	for _, r := range body {
		if old, err := core.DecodeBind(r); err == nil && old.EdgeID >= b.EdgeID {
			b.EdgeID = old.EdgeID + 1
		}
	}
	rec, err := core.EncodeBind(b)
	if err != nil {
		return nil, err
	}
	var in []core.Word
	for _, r := range head {
		ws, _ := core.EncodeRecord(r)
		in = append(in, ws...)
	}
	in = append(in, core.TargetWord(core.KindMeta, core.MetaBody))
	for _, r := range append(body, rec) {
		ws, _ := core.EncodeRecord(r)
		in = append(in, ws...)
	}
	return core.PutWords(core.WrapDocument(append(in, core.TargetWord(core.KindMeta, core.MetaEnd)))), nil
}
