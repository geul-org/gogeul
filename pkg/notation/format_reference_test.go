//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestFormatReference — #n, Qn, Pn, document addresses with and without version, ISSUER-POINTER, an Issuer document's LOCAL; a Quantifier or another codec is false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatReference(t *testing.T) {
	doc, _ := core.EncodeDocPointer([]core.Word{0, 0}, []core.Word{0, 0, 0, 0x806}, []core.Word{0, 2})
	latest, _ := core.EncodeDocPointer([]core.Word{0, 0}, []core.Word{0, 0, 0, 0x806}, []core.Word{0, 0})
	is, _ := core.EncodeIssuerPointer([]core.Word{1, 0})
	ext := core.Record{Kind: core.KindReference, Codec: core.CodecExternalIssuer2Word, Payload: []core.Word{0, 0, 1, 0, 0, 1, 2}}
	cases := map[string]core.Record{
		"#7":                              core.EncodeLocalRef(7),
		"Q42":                             {Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{42}},
		"0x00000000/0x0000000000000806@2": doc,
		"0x00000000/0x0000000000000806":   latest,
		"ISSUER-POINTER(0x00010000)":      is,
		"0x00000000/0x0001000000000001#2": ext,
	}
	for want, r := range cases {
		if got, ok := formatReference(r); !ok || got != want {
			t.Errorf("%s: got %s %v", want, got, ok)
		}
	}
	for _, r := range []core.Record{
		{Kind: core.KindReference, Codec: core.CodecLocal + 1, Payload: []core.Word{1}},
		{Kind: core.KindReference, Codec: core.CodecGeul, Payload: []core.Word{1}},
	} {
		if _, ok := formatReference(r); ok {
			t.Errorf("%+v: want false", r)
		}
	}
}
