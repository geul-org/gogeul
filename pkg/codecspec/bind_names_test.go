//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestBindNames — LOCAL Subjects get their last name; other Subjects and records are skipped
package codecspec

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestBindNames(t *testing.T) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	other := core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: slices.Concat(enc(core.EncodeLocalRef(3)), []core.Word{0x02F7, 0, 0, 0, 0, 0, 1, 0}, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("x"))))}
	n := bindNames([]core.Record{nameOf(1, 5, "a"), nameOf(2, 5, "b"), other, core.EncodeUintValue(1)})
	if len(n) != 1 || n[5] != "b" {
		t.Fatalf("got %v", n)
	}
}
