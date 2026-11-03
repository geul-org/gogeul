//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeNode — DecodeNode rejects non-node records, truncation, reserved bits, role mismatch, bad slots and leftovers
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestDecodeNode feeds DecodeNode malformed records: a Value record, a LINK
// codec, an empty payload, a missing header, reserved header bits, a rule whose
// header role disagrees with its codec, each truncated or invalid slot, a
// missing SPAN and a trailing word.
func TestDecodeNode(t *testing.T) {
	rec := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindEdge, Codec: c, Payload: p}
	}
	flag := func(f NodeFlags, rest ...core.Word) core.Record {
		return rec(CodecCaseDeclare, append([]core.Word{0x0051, 1, core.Word(f)}, rest...)...)
	}
	bad := []core.Record{
		{Kind: core.KindValue},
		rec(CodecAttacks),
		rec(CodecCaseDeclare),
		rec(CodecCaseDeclare, 1),
		rec(CodecCaseDeclare, 0x0051, 1),
		rec(CodecCaseDeclare, 0x0051, 1, 0x1000, 0, 0),
		rec(CodecCounterDeclare, 0x0051, 1, 0, 0, 0),
		flag(NodeFlagFunc),
		flag(NodeFlagFunc, 0x0051, 1),
		flag(NodeFlagQual),
		flag(NodeFlagQual, 0x0421, 1),
		flag(NodeFlagQual, 0x1000),
		flag(NodeFlagQual, 0x0900, 0, 0, 2),
		flag(NodeFlagCheck),
		flag(NodeFlagClaim),
		flag(NodeFlagClaim, 0x0051, 1),
		flag(NodeFlagClaim, 0x0051, 1, 2),
		flag(NodeFlagClaim, 0x0051, 1, 2, 0x0051, 3),
		rec(CodecCaseDeclare, 0x0051, 1, 0x0080|core.Word(NodeFlagClaim), 0x0051, 1, 2),
		rec(CodecCaseDeclare, 0x0051, 1, 0x0080|core.Word(NodeFlagClaim), 0x0051, 1, 2, 0x0421),
		rec(CodecCaseDeclare, 0x0051, 1, 0x0080|core.Word(NodeFlagClaim), 0x0051, 1, 2, 0x0840),
		rec(CodecCaseDeclare, 0x0051, 1, 0x0080|core.Word(NodeFlagClaim), 0x0051, 1, 2, 0x0982, 1, 0, 0, 0, 0, 0),
		rec(CodecCaseDeclare, 0x0051, 1, 0x0100|core.Word(NodeFlagClaim), 0x0051, 1, 2, 0),
		rec(CodecCaseDeclare, 0x0051, 1, 0x0180|core.Word(NodeFlagClaim), 0x0051, 1, 2, 0, 0, 0),
		flag(NodeFlagVerdict, 0x0840),
		flag(NodeFlagWith),
		flag(NodeFlagWith, 0x0421, 5, 0x0051, 1),
		flag(0, 0),
		flag(0, 0, 0, 0),
	}
	for i, r := range bad {
		if _, err := DecodeNode(r, testWire); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
