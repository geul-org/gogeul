//ff:func feature=inspect type=decoder control=sequence
//ff:what TestHeadParts — a direct codec has only the TAG; an ISSUER codec adds Issuer and format; a long record adds the Length extension
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestHeadParts(t *testing.T) {
	entry := func(ws []core.Word) Entry {
		r, _, _ := core.DecodeRecord(ws)
		return Entry{Record: r, Words: ws}
	}
	edge, _ := toulmin.EncodeEdge(toulmin.EdgePacket{ID: 1, Prop: toulmin.PropAttacks, Src: 2, DstTID: 3}, testWire)
	if p := headParts(entry(edge)); len(p) != 1 || p[0].Name != "TAG" {
		t.Fatalf("direct: %+v", p)
	}
	text, _ := toulmin.EncodeText(toulmin.TextPacket{TID: 1, Bytes: make([]byte, 40)}, testWire)
	p := headParts(entry(text))
	if len(p) != 4 || p[1].Name != "LENGTH" || p[2].Name != "ISSUER" || p[3].Name != "FORMAT" || p[3].Words[1] != 0xC001 {
		t.Fatalf("issuer + length: %+v", p)
	}
}
