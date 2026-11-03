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
	if p := headParts(entry(text)); len(p) != 2 || p[1].Name != "LENGTH" {
		t.Fatalf("direct + length: %+v", p)
	}
	ws, _ := core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{0, 7}, Format: 0xC001, Payload: make([]core.Word, 20)})
	p := headParts(entry(ws))
	if len(p) != 4 || p[1].Name != "LENGTH" || p[2].Name != "ISSUER" || p[3].Name != "FORMAT" || p[3].Words[1] != 0xC001 {
		t.Fatalf("issuer + length: %+v", p)
	}
}
