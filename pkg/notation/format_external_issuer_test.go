//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatExternalIssuer — a 2-word Issuer's LOCAL; a payload without a target and a non-minimal target are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatExternalIssuer(t *testing.T) {
	r := core.Record{Kind: core.KindReference, Codec: core.CodecExternalIssuer2Word, Payload: []core.Word{1, 0, 0x5100, 0, 0, 42, 3}}
	if s, ok := formatExternalIssuer(r); !ok || s != "0x00010000/0x510000000000002a#3" {
		t.Fatal(s)
	}
	r.Payload = r.Payload[:6]
	if _, ok := formatExternalIssuer(r); ok {
		t.Fatal("no target: want false")
	}
	r.Payload = []core.Word{1, 0, 0x5100, 0, 0, 42, 0, 3}
	if _, ok := formatExternalIssuer(r); ok {
		t.Fatal("non-minimal: want false")
	}
}
