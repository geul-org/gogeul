//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestWalkerField — raw words with ENUM names, the rest of the payload, records with their Value, slot markers as Marker where named, CODEC mismatch, unnamed markers and short input marked bad
package inspect

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestWalkerField(t *testing.T) {
	val, _ := core.EncodeRecord(core.EncodeUintValue(9))
	w := &walker{ws: append([]core.Word{1, 2}, val...)}
	w.field(codecspec.Field{Name: "code", Width: 1, Enums: map[string]string{"0x0001": "one"}})
	w.field(codecspec.Field{Name: "other", Width: 1})
	w.field(codecspec.Field{Name: "v", Codecs: []core.Word{0x0050, 0x040F}})
	if w.bad || w.out[0].Enum != "one" || w.out[1].Enum != "" || w.out[2].Data != (Value{Codec: core.CodecUnsignedInteger, Uint: 9}) {
		t.Fatalf("got %+v bad %v", w.out, w.bad)
	}
	r := &walker{ws: []core.Word{5, 6, 7}, pos: 1}
	r.field(codecspec.Field{Name: "rest", Width: int(core.FieldRest)})
	if r.bad || r.pos != 3 || len(r.out[0].Words) != 2 {
		t.Fatalf("rest: %+v", r)
	}
	r.field(codecspec.Field{Name: "empty", Width: int(core.FieldRest)})
	if r.bad || len(r.out) != 1 {
		t.Fatalf("empty rest: %+v", r)
	}
	m := &walker{ws: []core.Word{core.UnknownValue, core.NoValue}}
	m.field(codecspec.Field{Name: "object", Codecs: []core.Word{0x000F, core.UnknownValue, core.NoValue}})
	m.field(codecspec.Field{Name: "object", Codecs: []core.Word{0x000F, core.UnknownValue, core.NoValue}})
	if m.bad || m.out[0].Data != (Marker{Number: 1, Name: "UNKNOWN-VALUE"}) || m.out[1].Data != (Marker{Number: 2, Name: "NO-VALUE"}) {
		t.Fatalf("markers: %+v bad %v", m.out, m.bad)
	}
	for i, f := range []codecspec.Field{{Codecs: []core.Word{0x000F}}, {}} {
		w := &walker{ws: []core.Word{core.Absent}}
		w.field(f)
		if !w.bad {
			t.Fatalf("marker %d not named: want bad", i)
		}
	}
	for i, f := range []codecspec.Field{{Width: 3}, {Codecs: []core.Word{0x0980}}} {
		w := &walker{ws: val}
		w.field(f)
		if !w.bad {
			t.Fatalf("bad %d: want bad", i)
		}
	}
}
