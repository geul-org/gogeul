//ff:func feature=delta type=decoder control=sequence
//ff:what TestSplitSnapshot — HEAD and Edge Ids in order; odd bytes, a broken document and a non-Edge BODY record fail
package delta

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSplitSnapshot(t *testing.T) {
	s, err := splitSnapshot(encodeDoc([]core.Record{core.EncodeSpecVersion(2)}, []core.Record{bindOf(5, 9, "a"), bindOf(2, 9, "b")}))
	if err != nil || len(s.head) != 1 || !slices.Equal(s.ids, []uint64{5, 2}) {
		t.Fatalf("%+v %v", s, err)
	}
	if _, err := splitSnapshot([]byte{1}); err == nil {
		t.Fatal("odd: want error")
	}
	if _, err := splitSnapshot(core.PutWords([]core.Word{0x0051, 0x0001})); err == nil {
		t.Fatal("broken: want error")
	}
	if _, err := splitSnapshot(encodeDoc(nil, []core.Record{core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("ab"))})); err == nil {
		t.Fatal("non-Edge: want error")
	}
}
