//ff:func feature=resolve type=rule control=iteration dimension=1
//ff:what TestApplyListRecord — PUBLICATION adds, DEPRECATION marks, BIND passes; other documents, bad layouts and foreign records fail
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestApplyListRecord(t *testing.T) {
	doc := []core.Word{0xffff, 0, 0, 1}
	lg := &versionLog{}
	pub := core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: []core.Word{0, 0}, Doc: doc, Version: 1})
	dep := core.EncodePublication(core.Publication{Codec: core.EdgeDeprecation, EdgeID: 2, Issuer: []core.Word{0, 0}, Doc: doc, Version: 1})
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	bind := core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: enc(core.EncodeLocalRef(3))}
	for i, r := range []core.Record{pub, dep, bind} {
		if err := applyListRecord(lg, r, "0xffff000000000001"); err != nil {
			t.Fatalf("ok %d: %v", i, err)
		}
	}
	if len(lg.Entries) != 1 || !lg.Entries[0].Deprecated {
		t.Fatalf("got %+v", lg.Entries)
	}
	other := core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 4, Issuer: []core.Word{0, 0}, Doc: []core.Word{0xffff, 0, 0, 2}, Version: 2})
	for i, r := range []core.Record{other, core.EncodeUintValue(1), {Kind: core.KindEdge, Codec: core.EdgePublication}} {
		if err := applyListRecord(lg, r, "0xffff000000000001"); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
