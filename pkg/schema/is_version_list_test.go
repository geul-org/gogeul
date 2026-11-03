//ff:func feature=schema type=rule control=sequence
//ff:what TestIsVersionList — no DOC-ID and a leading PUBLICATION; a DOC-ID, another first record or an empty BODY are not
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestIsVersionList(t *testing.T) {
	pub := recOf(edgeOf(core.EdgePublication, 1, append([]core.Word{0x03D8, 0, 0xffff, 0, 0, 2, 0, 1}, make([]core.Word, 16)...)...))
	id := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: []core.Word{0, 0, 0, 0, 2}})
	if !isVersionList(docParts{Body: []rec{pub}}) || isVersionList(docParts{Head: []rec{id}, Body: []rec{pub}}) || isVersionList(docParts{Body: []rec{recOf(edgeOf(0, 1))}}) || isVersionList(docParts{}) {
		t.Fatal("wrong")
	}
}
