//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeSource — DecodeSource rejects non-SOURCE records, a bad width word, a bad Value and leftovers
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestDecodeSource feeds DecodeSource malformed records.
func TestDecodeSource(t *testing.T) {
	src := func(p ...core.Word) core.Record {
		return core.Record{Kind: core.KindEdge, Codec: CodecSource, Payload: p}
	}
	bad := []core.Record{
		{Kind: core.KindValue},
		{Kind: core.KindEdge, Codec: CodecText},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{0, 0}, Format: 0xC000},
		src(),
		src(0x0051, 1),
		src(0x0051, 1, 0x0001),
		src(0x0051, 1, 0),
		src(0x0051, 1, 0, 0x0621, 1),
		src(0x0051, 1, 0, 0x0A40, 0),
	}
	for i, r := range bad {
		if _, _, err := DecodeSource(r); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
