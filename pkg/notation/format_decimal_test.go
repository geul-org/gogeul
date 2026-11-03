//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestFormatDecimal — SIGNED-DECIMAL bare, UNSIGNED-DECIMAL by name; broken payloads are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatDecimal(t *testing.T) {
	sd, _ := core.EncodeDecimalValue(true, -1, -15)
	ud, _ := core.EncodeDecimalValue(false, 0, 7)
	if s, ok := formatDecimal(sd); !ok || s != "-1.5" {
		t.Fatal(s)
	}
	if s, ok := formatDecimal(ud); !ok || s != "UNSIGNED-DECIMAL(7)" {
		t.Fatal(s)
	}
	for _, c := range []uint16{core.CodecSignedDecimal, core.CodecUnsignedDecimal} {
		if _, ok := formatDecimal(core.Record{Kind: core.KindValue, Codec: c, Payload: []core.Word{0}}); ok {
			t.Fatal("broken: want false")
		}
	}
}
