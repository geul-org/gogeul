//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckBitVector — grammar examples pass (empty, `101`); short counts, wrong word counts and nonzero padding are V2
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckBitVector(t *testing.T) {
	bv := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindValue, Codec: core.CodecBitVector, Payload: p})
	}
	for _, r := range []rec{bv(0, 0), bv(0, 3, 0xa000), bv(0, 16, 0xffff), bv(0, 17, 0xffff, 0x8000)} {
		if fs := checkBitVector(r); fs != nil {
			t.Fatalf("%04X: %v", r.R.Payload, fs)
		}
	}
	for _, r := range []rec{bv(), bv(0), bv(0, 0, 1), bv(0, 3), bv(0, 3, 0xa000, 0), bv(0, 3, 0xb000), bv(0, 17, 0xffff, 0x8001)} {
		if got := rulesOf(checkBitVector(r)); got != "V2" {
			t.Fatalf("%04X: %q", r.R.Payload, got)
		}
	}
}
