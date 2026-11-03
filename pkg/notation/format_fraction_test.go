//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatFraction — FRACTION(23.45, 1256.1); a broken FRACTION and one with an exponent-0 DECIMAL component are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatFraction(t *testing.T) {
	f := core.Record{Kind: core.KindValue, Codec: core.CodecFraction, Payload: []core.Word{0x0483, 0xFFFF, 0xFFFE, 0x0929, 0x0493, 0xFFFF, 0xFFFF, 0x3111}}
	if s, ok := formatFraction(f); !ok || s != "FRACTION(23.45, 1256.1)" {
		t.Fatalf("got %s %v", s, ok)
	}
	if _, ok := formatFraction(core.Record{Kind: core.KindValue, Codec: core.CodecFraction}); ok {
		t.Fatal("broken: want false")
	}
	zero := core.Record{Kind: core.KindValue, Codec: core.CodecFraction, Payload: []core.Word{0x0483, 0, 0, 5, 0x0421, 3}}
	if _, ok := formatFraction(zero); ok {
		t.Fatal("exponent 0: want false")
	}
}
