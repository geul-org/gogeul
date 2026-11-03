//ff:func feature=cli type=helper control=sequence
//ff:what TestReferrer — the Edge referring to a gone ID, in the BODY or the HEAD; none when nothing refers
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/notation"
)

func TestReferrer(t *testing.T) {
	ref, _ := core.EncodeRecord(core.EncodeLocalRef(2))
	name, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("a")))
	d := draft{body: []core.Record{edgeRecord(3, notation.Edge{Codec: core.EdgeBind, Rest: append(ref, name...)})}}
	if by, ok := referrer(d, []uint64{2}); !ok || by != 3 {
		t.Fatalf("%d %v", by, ok)
	}
	zero, _ := core.EncodeRecord(core.EncodeLocalRef(0))
	q1 := []core.Word{0x0141, 1}
	p813 := []core.Word{0x0191, 813}
	head := draft{head: []core.Record{
		edgeRecord(2, notation.Edge{Codec: core.EdgeSource, Rest: append(zero, q1...)}),
		edgeRecord(3, notation.Edge{Codec: core.EdgeQualifier, Rest: append(append(append([]core.Word{}, p813...), ref...), q1...)}),
	}}
	if by, ok := referrer(head, []uint64{2}); !ok || by != 3 {
		t.Fatalf("HEAD: %d %v", by, ok)
	}
	if _, ok := referrer(d, []uint64{7}); ok {
		t.Fatal("none: want false")
	}
}
