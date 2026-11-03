//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what TestEdgeID — an Edge's LOCAL Edge Id and the rest; a Meta, an Edge without a LOCAL first and Edge Id 0 fail
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEdgeID(t *testing.T) {
	if id, rest, ok := edgeID(bindOf(7, 2, "a")); !ok || id != 7 || len(rest) == 0 {
		t.Fatalf("%d %v", id, ok)
	}
	for _, r := range []core.Record{core.EncodeSpecVersion(2), {Kind: core.KindEdge, Codec: core.EdgeBind, Payload: []core.Word{0x0AC0, 0x6162}}, bindOf(0, 2, "a")} {
		if _, _, ok := edgeID(r); ok {
			t.Fatalf("%+v: want false", r)
		}
	}
}
