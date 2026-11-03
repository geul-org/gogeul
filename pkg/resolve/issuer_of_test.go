//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestIssuerOf — an ISSUER-POINTER's canonical 2/4/16-word Issuer as hex; zero-padded, 1-word and other Reference codes fail
package resolve

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestIssuerOf(t *testing.T) {
	ref := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindReference, Codec: c, Payload: p}
	}
	if s, ok := issuerOf(ref(core.CodecIssuerPointer, 0, 0)); !ok || s != "0x00000000" {
		t.Fatalf("2 words: %q %v", s, ok)
	}
	if s, ok := issuerOf(ref(core.CodecIssuerPointer, 1, 0x1234)); !ok || s != "0x00011234" {
		t.Fatalf("2 words: %q %v", s, ok)
	}
	if s, ok := issuerOf(ref(core.CodecIssuerPointer, append(make([]core.Word, 15), 1)...)); ok {
		t.Fatalf("16 words of a 2-word value: %q", s)
	}
	if s, ok := issuerOf(ref(core.CodecIssuerPointer, append([]core.Word{1}, make([]core.Word, 15)...)...)); !ok || s != "0x0001"+strings.Repeat("0", 60) {
		t.Fatalf("16 words: %q %v", s, ok)
	}
	bad := []core.Record{
		ref(core.CodecExternalIssuer2Word, 0, 0, 0, 0, 0, 0, 0),
		ref(core.CodecDocPointer, 0, 0, 0xffff, 0, 0, 1, 0, 0),
		ref(core.CodecIssuerPointer, 0),
		ref(core.CodecIssuerPointer, 0, 0, 0),
		{Kind: core.KindValue, Codec: core.CodecIssuerPointer, Payload: []core.Word{0, 0}},
	}
	for i, r := range bad {
		if _, ok := issuerOf(r); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
}
