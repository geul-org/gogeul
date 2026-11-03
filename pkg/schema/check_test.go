//ff:func feature=schema type=rule control=sequence
//ff:what TestCheck — a valid document passes, an unassigned code too (§1.1 112); odd bytes, a Warn (M4), Strict and Codecs change the report
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheck(t *testing.T) {
	ok := docOf(nil, []core.Record{core.EncodeUintValue(5)})
	if r := Check(ok, Options{}); r.Verdict != Pass || len(r.Facts) != 0 {
		t.Fatalf("ok: %+v", r)
	}
	if r := Check(docOf(nil, []core.Record{{Kind: core.KindValue, Codec: 0x22}}), Options{Strict: true}); r.Verdict != Pass {
		t.Fatalf("unassigned code reported (§1.1 112): %+v", r)
	}
	if r := Check(append(ok, 0), Options{}); r.Verdict != Fail || rulesOf(r.Facts) != "F1" {
		t.Fatalf("odd: %+v", r)
	}
	unknown := docOf([]core.Record{{Kind: core.KindMeta, Codec: core.MetaSpecVersion, Payload: []core.Word{3}}}, []core.Record{core.EncodeUintValue(5)})
	if r := Check(unknown, Options{}); r.Verdict != Review {
		t.Fatalf("unknown SPEC-VERSION: %+v", r)
	}
	if r := Check(unknown, Options{Strict: true}); r.Verdict != Fail {
		t.Fatalf("strict: %+v", r)
	}
	if r := Check(ok, Options{Codecs: []Codec{fakeCodec{codec: core.CodecUnsignedInteger}}}); rulesOf(r.Facts) != "X1" {
		t.Fatalf("codecs: %+v", r)
	}
}
