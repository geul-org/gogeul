//ff:func feature=inspect type=decoder control=sequence
//ff:what TestFractionText — 2/4 stays 2/4 (not reduced); a broken FRACTION, a non-minimal numerator and a non-minimal denominator fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFractionText(t *testing.T) {
	f := func(p ...core.Word) core.Record {
		return core.Record{Kind: core.KindValue, Codec: core.CodecFraction, Payload: p}
	}
	if s, err := fractionText(f(0x0611, 2, 0x0621, 4)); err != nil || s != "2/4" {
		t.Fatalf("2/4: %q %v", s, err)
	}
	if _, err := fractionText(f()); err == nil {
		t.Fatal("broken: want error")
	}
	if _, err := fractionText(f(0x0612, 0, 2, 0x0621, 4)); err == nil {
		t.Fatal("non-minimal numerator: want error")
	}
	if _, err := fractionText(f(0x0611, 2, 0x0622, 0, 4)); err == nil {
		t.Fatal("non-minimal denominator: want error")
	}
}
