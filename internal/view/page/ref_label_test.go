//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestRefLabel — Reference fields in hex, a lone TAG empty, another payload as one run
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRefLabel(t *testing.T) {
	ext := core.Record{Kind: core.KindReference, Codec: core.CodecExternalIssuer2Word, Payload: []core.Word{0, 0, 0xffff, 0, 0, 1, 0}}
	c, _, _ := core.DecodeRecord([]core.Word{0x0980})
	for want, r := range map[string]core.Record{
		"0x00000000 0xffff000000000001 0x0000": ext,
		"0x0007":                               core.EncodeLocalRef(7),
		"":                                     c,
		"0x00010002":                           {Kind: core.KindEdge, Codec: 5, Payload: []core.Word{1, 2}},
	} {
		if got := refLabel(r); got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}
