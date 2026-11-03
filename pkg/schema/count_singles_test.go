//ff:func feature=schema type=rule control=sequence
//ff:what TestCountSingles — a repeated single Meta is reported once per repeat by its rule (M2, M8); HEAD Edges sharing a Meta code and MESSAGEs of different languages are not
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCountSingles(t *testing.T) {
	m := func(c uint16) rec { return recOf(core.Record{Kind: core.KindMeta, Codec: c}) }
	if fs := countSingles([]rec{m(core.MetaDocID), m(core.MetaEffectiveFrom), m(core.MetaTag), m(core.MetaTag), m(core.MetaDocID)}); rulesOf(fs) != "M2" {
		t.Fatalf("got %+v", fs)
	}
	if fs := countSingles([]rec{m(core.MetaAuthor), m(core.MetaRecordedAt), m(core.MetaAuthor), m(core.MetaRecordedAt)}); rulesOf(fs) != "M8,M8" {
		t.Fatalf("got %+v", fs)
	}
	src := recOf(edgeOf(core.EdgeSource, 3, 0x0051, 0, 0x0141, 42))
	if fs := countSingles([]rec{src, src, m(core.MetaPrivacy)}); fs != nil {
		t.Fatalf("HEAD Edges (SOURCE = 0x20 = PRIVACY's code): %+v", fs)
	}
	msg := func(lang core.Word) rec {
		v := core.Record{Kind: core.KindValue, Codec: core.CodecUTF8LanguageTextEven, Payload: []core.Word{0x07A1, lang, 0x6869}}
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaMessage, Payload: enc(v)})
	}
	if fs := countSingles([]rec{msg(1), msg(2), msg(1), m(core.MetaMessage), m(core.MetaMessage)}); rulesOf(fs) != "M8" {
		t.Fatalf("messages: %+v", fs)
	}
}
