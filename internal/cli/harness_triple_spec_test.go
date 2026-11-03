//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what harness.tripleSpec — test fixture: publish TRIPLE's codec spec (three record fields and a 1-word rank: normal, preferred) in the data folder
package cli

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// tripleSpec publishes the spec document of TRIPLE (0x0A) as geul.org's
// codec for its target word (resolve.Dir.PublishCodec).
func (h *harness) tripleSpec() {
	target := core.TargetWord(core.KindEdge, 0x0A)
	id, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, uint32(target))
	body := []core.Record{core.EncodeDeclare(core.Declare{EdgeID: 1, Target: target})}
	for i, w := range []core.Word{0, 0, 0, 1} {
		body = append(body, core.EncodeFieldDecl(core.FieldDecl{EdgeID: uint64(i + 2), Declare: 1, Width: w}))
	}
	for i, name := range []string{"normal", "preferred"} {
		body = append(body, core.EncodeEnumDecl(core.EnumDecl{EdgeID: uint64(6 + 2*i), Field: 5, Code: []core.Word{core.Word(i)}}))
		b, _ := core.EncodeBind(core.Bind{EdgeID: uint64(7 + 2*i), Subject: core.EncodeLocalRef(uint64(6 + 2*i)), Name: []byte(name)})
		body = append(body, b)
	}
	doc, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(core.SpecVersion), id}, body)
	if _, err := (resolve.Dir{Root: h.data}).PublishCodec(doc); err != nil {
		h.t.Fatal(err)
	}
}
