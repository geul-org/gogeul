//ff:func feature=inspect type=decoder control=sequence
//ff:what TestSpecParts — a declared record splits into Edge Id and fields, a rest field before fixed ones; no spec, a short or long payload do not fit
package inspect

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestSpecParts(t *testing.T) {
	r := core.EncodePublication(core.Publication{Codec: core.EdgeTombstone, EdgeID: 3, Issuer: []core.Word{0, 0}, Doc: []core.Word{0, 0, 0, 1}, Version: 1, Reason: 1})
	fs := tombstoneSpec()[core.TargetWord(core.KindEdge, core.EdgeTombstone)]
	ps, ok := specParts(Entry{Record: r, Spec: fs})
	if !ok || len(ps) != 3 || ps[0].Name != "ID" || ps[1].Name != "doc" || !ps[1].Record || ps[2].Name != "reason" || ps[2].Enum != "license" {
		t.Fatalf("got %+v %v", ps, ok)
	}
	doc := core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: []core.Word{0, 0x20, 0xffff, 0, 0, 1}}
	if ps, ok := specParts(Entry{Record: doc, Spec: []codecspec.Field{{Name: "issuer", Width: 0xFFFF}, {Name: "doc", Width: 4}}}); !ok || len(ps[0].Words) != 2 || len(ps[1].Words) != 4 {
		t.Fatalf("rest first: %+v %v", ps, ok)
	}
	if _, ok := specParts(Entry{Record: doc, Spec: []codecspec.Field{{Width: 0xFFFF}, {}}}); ok {
		t.Fatal("record after rest: want false")
	}
	if _, ok := specParts(Entry{Record: r}); ok {
		t.Fatal("no spec: want false")
	}
	if _, ok := specParts(Entry{Record: r, Spec: fs[:1]}); ok {
		t.Fatal("left over: want false")
	}
	if _, ok := specParts(Entry{Record: r, Spec: append(fs, codecspec.Field{Width: 1})}); ok {
		t.Fatal("short: want false")
	}
	v := core.Record{Kind: core.KindMeta, Codec: core.MetaRecordedAt, Payload: func() []core.Word { ws, _ := core.EncodeRecord(core.EncodePosixSeconds(5)); return ws }()}
	if ps, ok := specParts(Entry{Record: v, Spec: []codecspec.Field{{Name: "at"}}}); !ok || len(ps) != 1 || ps[0].Data == nil {
		t.Fatalf("meta: %+v %v", ps, ok)
	}
}
