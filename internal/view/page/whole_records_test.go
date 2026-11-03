//ff:func feature=page type=helper control=sequence
//ff:what TestWholeRecords — a text record and two records are whole; one word, a partial record and an Edge reserved codecs and slot markers are not
package page

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestWholeRecords(t *testing.T) {
	a, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("v1")))
	b, _ := core.EncodeRecord(core.EncodeLocalRef(3))
	e, _ := core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: 0x03, Payload: []core.Word{0x0051, 1}})
	if !wholeRecords(a) || !wholeRecords(slices.Concat(a, b)) {
		t.Fatal("whole: want true")
	}
	if wholeRecords([]core.Word{0x0051}) || wholeRecords(a[:1]) || wholeRecords(e) || wholeRecords([]core.Word{0xFFFF, 0}) || wholeRecords([]core.Word{0, 2}) || wholeRecords([]core.Word{1, 2}) || wholeRecords(slices.Concat(b, []core.Word{core.NoValue})) {
		t.Fatal("not whole: want false")
	}
}
