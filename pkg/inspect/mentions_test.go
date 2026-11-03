//ff:func feature=inspect type=decoder control=sequence
//ff:what TestMentions — an Edge's mentions without its Edge Id; a Meta's LOCAL counts; no records mention nothing
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMentions(t *testing.T) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	q := core.Record{Kind: core.KindReference, Codec: core.CodecWikidataPID, Payload: []core.Word{580}}
	qual := core.Record{Kind: core.KindEdge, Codec: 0x0B, Payload: slices.Concat(enc(core.EncodeLocalRef(4)), enc(q), enc(core.EncodeLocalRef(2)), enc(core.EncodeLocalRef(3)))}
	if got := Mentions(newEntry(qual)); !slices.Equal(got, []uint64{2, 3}) {
		t.Fatal(got)
	}
	remove := core.EncodeDeltaRemove([]uint64{5})
	if got := Mentions(newEntry(remove)); len(got) != 0 && got[0] != 5 {
		t.Fatal(got)
	}
	if got := Mentions(newEntry(core.EncodeSpecVersion(2))); len(got) != 0 {
		t.Fatal(got)
	}
}
