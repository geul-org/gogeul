//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckValue — Value rules per direct code (new C numbers), VALUE-ISSUER skipped, empty codes V1
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckValue(t *testing.T) {
	v := func(c uint16, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindValue, Codec: c, Payload: p})
	}
	cases := []struct {
		r    rec
		want string
	}{
		{v(0x3D), ""},
		{v(0x3F), ""},
		{v(0x3C), "V1"},
		{v(0x00, 2), "V2"},
		{v(0x01, 0, 1), "V3"},
		{v(0x02, 0, 1), "V3"},
		{v(0x03, 1, 2, 3), "V2"},
		{v(0x09, 0), "V3"},
		{v(0x0B, 1), "V2"},
		{v(0x0D, 1), "V2"},
		{v(0x1A, 1, 2), "V2"},
		{v(0x1A, 7), ""},
		{v(0x1B, 1), "V2"},
		{v(0x1C, 1), "V2"},
		{v(0x1D, 1), "V2"},
		{v(0x1D, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0), ""},
		{v(0x1E), "V1"},
		{v(0x13), "V4"},
		{v(0x0F), "V4"},
		{v(0x0E, 0x3139, 0x3734), ""},
		{v(0x11), "V4"},
		{v(0x14, 0x6162), ""},
		{v(0x15), "V4"},
		{v(0x17), "V4"},
		{v(0x04, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0), ""},
		{v(0x05), "V8"},
		{v(0x06), "V9"},
		{v(0x07), "V10"},
		{v(0x07, 0x0411, 2, 0x0421, 3), ""},
		{v(0x30), "V1"},
		{v(0x18, 0x0561, 0x656E), ""},
	}
	for i, c := range cases {
		if got := rulesOf(checkValue(c.r)); got != c.want {
			t.Fatalf("case %d: %q", i, got)
		}
	}
}
