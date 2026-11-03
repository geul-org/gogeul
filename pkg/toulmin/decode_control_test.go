//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeControl — DecodeControl rejects non-control records, missing fields and leftovers
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestDecodeControl feeds DecodeControl records that are not framing controls
// (wrong kind, a foreign Issuer, a direct codec, a TEXT format) and controls
// with missing or extra payload words.
func TestDecodeControl(t *testing.T) {
	ctl := func(f uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindEdge, Codec: f, Payload: p}
	}
	bad := []core.Record{
		{Kind: core.KindValue},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{0, 0}, Format: 0xC002},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer4Word, Issuer: []core.Word{0, 1, 0, 0}, Format: 0xC002},
		{Kind: core.KindEdge, Codec: CodecAttacks},
		ctl(CodecText, 1),
		ctl(CodecSection),
		ctl(CodecSection, 1),
		ctl(CodecSection, 0x0051, 1),
		ctl(CodecProse, 0x0051, 1, 0),
		ctl(CodecCaseEnd, 0x0051, 1),
		ctl(CodecCaseEnd, 0x0051, 1, 0x0051, 2, 3),
	}
	for i, r := range bad {
		if _, err := DecodeControl(r, testWire); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
