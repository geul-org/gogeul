//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseHex — words from 0x and groups of 4 digits, either case; missing prefix, odd lengths and non-hex fail
package notation

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseHex(t *testing.T) {
	if ws, err := parseHex("0x0B86FFFF"); err != nil || !slices.Equal(ws, []core.Word{0x0B86, 0xFFFF}) {
		t.Fatalf("%v %v", ws, err)
	}
	for _, s := range []string{"0B86", "0x", "0x123", "0xzzzz"} {
		if _, err := parseHex(s); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
