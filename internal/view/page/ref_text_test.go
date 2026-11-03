//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestRefText — LOCAL, QID and PID by number; other codecs, Quantifiers, empty or long payloads and non-References are false
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRefText(t *testing.T) {
	ref := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindReference, Codec: c, Payload: p}
	}
	for want, r := range map[string]core.Record{
		"#9":       core.EncodeLocalRef(9),
		"Q42":      ref(core.CodecWikidataQID, 42),
		"P2860":    ref(core.CodecWikidataPID, 2860),
		"Q1234567": ref(core.CodecWikidataQID, 0x12, 0xD687),
	} {
		if got, ok := refText(r); !ok || got != want {
			t.Fatalf("got %q %v want %q", got, ok, want)
		}
	}
	for i, r := range []core.Record{
		ref(core.CodecGeul, 1), ref(core.CodecWikidataQID+1, 1), ref(core.CodecWikidataQID),
		ref(core.CodecWikidataQID, 1, 2, 3, 4, 5), {Kind: core.KindValue, Payload: []core.Word{1}},
	} {
		if _, ok := refText(r); ok {
			t.Fatalf("case %d: want false", i)
		}
	}
}
