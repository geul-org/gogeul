//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestTargets — each direct codec once and sorted, Reference Quantifier cleared; ISSUER records and Reference 0000 left out; the language registry when a payload holds a LANGUAGE-SHORT-CODE
package codecspec

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTargets(t *testing.T) {
	edge := core.Record{Kind: core.KindEdge, Codec: 0x0A}
	got := Targets([]core.Record{
		edge, edge,
		{Kind: core.KindReference, Codec: 0x0B},
		{Kind: core.KindReference, Codec: 0x03},
		{Kind: core.KindEdge, Codec: 0x7C, Issuer: []core.Word{1}},
	})
	want := []core.Word{core.TargetWord(core.KindReference, 0x08), core.TargetWord(core.KindEdge, 0x0A)}
	if !slices.Equal(got, want) {
		t.Fatalf("got %#x want %#x", got, want)
	}
	text := core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: []core.Word{0x0088, 1, 0x1234, languageTag, 0x0784}}
	if got := Targets([]core.Record{text}); !slices.Equal(got, []core.Word{languageCodec, core.TargetWord(core.KindEdge, 0x04)}) {
		t.Fatalf("got %#x", got)
	}
	if got := Targets(nil); len(got) != 0 {
		t.Fatalf("got %#x", got)
	}
}
