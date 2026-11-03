//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckReference — rules by C at the §7.3 boundaries: markers R1, pointers R5, schemes R2/R3/R4, CODEC-POINTER R5, REFERENCE-ISSUER nothing
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckReference(t *testing.T) {
	ref := func(codec uint16, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: codec, Payload: p})
	}
	cases := []struct {
		r    rec
		want string
	}{
		{ref(2, 1), "R1"},
		{ref(3, 0x40), "R5"},
		{ref(4, 0, 0, 0, 0, 0, 0, 1), "R5"},
		{ref(5, 1), ""},
		{ref(5, 0, 1), "R2"},
		{ref(0x14, 1), ""},
		{ref(0x1D, 1), ""},
		{ref(0x1E, 9, 1), "R3"},
		{ref(0x2C, 0, 0, 9, 1), "R3"},
		{ref(0x2D, 1), "R4"},
		{ref(0x36, 0, 1, 0, 0, 0, 0, 0, 0, 5), ""},
		{ref(0x37, 1), "R4"},
		{ref(0x3B, 1), "R4"},
		{ref(0x3C, 0, 0, 0, 0x980, 0, 0), ""},
		{ref(0x3C, 0, 0x980, 0, 0), "R5"},
	}
	for i, c := range cases {
		if got := rulesOf(checkReference(c.r, nil)); got != c.want {
			t.Fatalf("%d C 0x%02X: %q want %q", i, c.r.R.Codec, got, c.want)
		}
	}
	iss := recOf(core.Record{Kind: core.KindReference, Codec: core.CodecRefIssuer2Word, Issuer: []core.Word{0, 2}, Format: 1, Payload: []core.Word{0, 0}})
	if fs := checkReference(iss, nil); fs != nil {
		t.Fatalf("REFERENCE-ISSUER Payload is the Issuer's: %+v", fs)
	}
}
