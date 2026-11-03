//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestFormatReferenceBare — a DOC-POINTER or ISSUER-POINTER whose payload does not decode is false, not a panic; every 1-WORD record (a raw code such as ENUM 0x0038) formats
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatReferenceBare(t *testing.T) {
	for _, r := range []core.Record{
		{Kind: core.KindReference, Codec: core.CodecDocPointer},                                      // raw 0x0040
		{Kind: core.KindReference, Codec: core.CodecDocPointer, Payload: []core.Word{0}},             // raw 0x0041 0x0000
		{Kind: core.KindReference, Codec: core.CodecIssuerPointer},                                   // raw 0x0030
		{Kind: core.KindReference, Codec: core.CodecIssuerPointer, Payload: []core.Word{0, 0, 0, 1}}, // Issuer 0x0000000000000001 not canonical
	} {
		if s, ok := formatReference(r); ok || s != "" {
			t.Errorf("%+v: got %q %v, want false", r, s, ok)
		}
	}
	for v := 0; v < 0x10000; v++ {
		r, _, err := core.DecodeRecord([]core.Word{core.Word(v)})
		if err != nil {
			continue
		}
		if s := formatRecord(r); s == "" {
			t.Errorf("%04x: empty token", v)
		}
	}
	if got := formatRecord(core.Record{Kind: core.KindReference, Codec: core.CodecDocPointer}); got != "0x0040" {
		t.Errorf("bare DOC-POINTER: got %s, want raw 0x0040", got)
	}
}
