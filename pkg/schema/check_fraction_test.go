//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckFraction — 2/3 passes; a wrong layout and a zero denominator are V10
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckFraction(t *testing.T) {
	v := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindValue, Codec: core.CodecFraction, Payload: p})
	}
	if fs := checkFraction(v(0x0611, 2, 0x0621, 3)); fs != nil {
		t.Fatalf("2/3: %+v", fs)
	}
	if rulesOf(checkFraction(v(0x0611, 2))) != "V10" {
		t.Fatal("layout: want V10")
	}
	if rulesOf(checkFraction(v(0x0611, 2, 0x0621, 0))) != "V10" {
		t.Fatal("zero denominator: want V10")
	}
}
