//ff:func feature=inspect type=decoder control=sequence
//ff:what TestBodyEntry — Values, profile gating, non-TOULMIN records, unnamed codecs and the lossless check
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestBodyEntry(t *testing.T) {
	w := toulmin.Wire{Span: 1}
	if e := bodyEntry(core.EncodeUintValue(2), false, &w); e.Data != (Value{Codec: core.CodecUnsignedInteger, Uint: 2}) {
		t.Fatalf("value: %+v", e)
	}
	link, _ := toulmin.EncodeEdge(toulmin.EdgePacket{ID: 5, Prop: toulmin.PropAttacks, Src: 1, DstTID: 2}, w)
	r, _, _ := core.DecodeRecord(link)
	if e := bodyEntry(r, false, &w); e.Data != nil || e.Name != "" {
		t.Fatalf("no profile: %+v", e)
	}
	if e := bodyEntry(core.Record{Kind: core.KindMeta, Codec: core.MetaPrivacy}, true, &w); e.Data != nil {
		t.Fatalf("not an edge: %+v", e)
	}
	if e := bodyEntry(core.Record{Kind: core.KindEdge, Codec: 0x66}, true, &w); e.Name != "" || e.Data != nil {
		t.Fatalf("unnamed: %+v", e)
	}
	if e := bodyEntry(r, true, &w); e.Name != "ATTACKS" || e.ID != 5 || len(e.Refs) != 2 {
		t.Fatalf("link: %+v", e)
	}
	// A non-trimmed qualifier (80e-2) decodes but re-encodes as 8e-1: not lossless.
	q := toulmin.VFix(8000)
	node, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 1, Type: toulmin.NodeCase, Flags: toulmin.NodeFlagQual, Qual: &q}, w)
	rn, _, _ := core.DecodeRecord(node)
	copy(rn.Payload[3:7], []core.Word{0x0683, 0xFFFF, 0xFFFE, 80})
	if _, err := toulmin.DecodeNode(rn, w); err != nil {
		t.Fatalf("crafted node must decode: %v", err)
	}
	if e := bodyEntry(rn, true, &w); e.Name != "CASE-DECLARE" || e.Data != nil {
		t.Fatalf("lossy: %+v", e)
	}
}
