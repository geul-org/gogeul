//ff:func feature=inspect type=decoder control=sequence
//ff:what TestMaxLocalID — the largest of Edge Ids and LOCAL record fields; a document with no LOCAL gives 0
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMaxLocalID(t *testing.T) {
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
	// BIND #3: #9 = "x" — the subject LOCAL 9 is larger than any Edge Id.
	d, err := Decode(doc(bind(3, 9), bind(5, 2)))
	if err != nil || MaxLocalID(d) != 9 {
		t.Fatalf("got %d %v", MaxLocalID(d), err)
	}
	if empty, _ := Decode(doc()); MaxLocalID(empty) != 0 {
		t.Fatal("empty: want 0")
	}
}
