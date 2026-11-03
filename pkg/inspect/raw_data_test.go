//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestRawData — raw words read as an UNSIGNED-INTEGER; no CODEC, two, a Reference target, a wildcard and an undecodable payload give nil
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRawData(t *testing.T) {
	u := core.TargetWord(core.KindValue, core.CodecUnsignedInteger)
	if got := rawData([]core.Word{u}, []core.Word{2}); got != (Value{Codec: core.CodecUnsignedInteger, Uint: 2}) {
		t.Fatalf("uint: %+v", got)
	}
	for i, c := range [][]core.Word{nil, {u, u}, {core.TargetWord(core.KindReference, 0x05)}, {0x060F}} {
		if got := rawData(c, []core.Word{2}); got != nil {
			t.Fatalf("case %d: %+v", i, got)
		}
	}
	if got := rawData([]core.Word{u}, nil); got != nil {
		t.Fatalf("empty: %+v", got)
	}
}
