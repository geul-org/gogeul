//ff:func feature=inspect type=decoder control=sequence
//ff:what TestLocalIDs — Edge Ids and mentioned LOCALs in order; a document with no LOCAL gives none
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLocalIDs(t *testing.T) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	doc := func(body ...core.Record) []byte {
		in := append(enc(core.EncodeSpecVersion(core.SpecVersion)), 0x1C30)
		for _, r := range body {
			in = append(in, enc(r)...)
		}
		return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
	}
	bind := func(id, subject uint64) core.Record {
		p := slices.Concat(enc(core.EncodeLocalRef(id)), enc(core.EncodeLocalRef(subject)), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("x"))))
		return core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p}
	}
	d, err := Decode(doc(bind(3, 9), bind(5, 3)))
	if got := LocalIDs(d); err != nil || !slices.Equal(got, []uint64{3, 9, 5, 3}) {
		t.Fatalf("got %v %v", got, err)
	}
	if empty, _ := Decode(doc()); len(LocalIDs(empty)) != 0 {
		t.Fatal("empty: want none")
	}
}
