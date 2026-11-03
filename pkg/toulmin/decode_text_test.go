//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeText — DecodeText rejects non-TEXT records, missing ID/Value, non-text Values and leftovers
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestDecodeText feeds DecodeText malformed records.
func TestDecodeText(t *testing.T) {
	txt := func(p ...core.Word) core.Record {
		return core.Record{Kind: core.KindEdge, Codec: CodecText, Payload: p}
	}
	bad := []core.Record{
		{Kind: core.KindValue},
		{Kind: core.KindEdge, Codec: CodecAttacks},
		txt(),
		txt(0x0051, 1),
		txt(0x0051, 1, 0x0621, 1),
		txt(0x0051, 1, 0x0A40, 0),
	}
	for i, r := range bad {
		if _, err := DecodeText(r, testWire); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
