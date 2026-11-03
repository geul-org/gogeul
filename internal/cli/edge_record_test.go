//ff:func feature=cli type=helper control=sequence
//ff:what TestEdgeRecord — the Edge Id leads the payload
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/notation"
)

func TestEdgeRecord(t *testing.T) {
	r := edgeRecord(5, notation.Edge{Codec: 0x04, Rest: []core.Word{0x0141, 1}})
	if r.Kind != core.KindEdge || r.Codec != 0x04 || r.Payload[0] != 0x0051 || r.Payload[1] != 5 || len(r.Payload) != 4 {
		t.Fatalf("%+v", r)
	}
}
