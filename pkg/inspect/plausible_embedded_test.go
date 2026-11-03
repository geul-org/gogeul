//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestPlausibleEmbedded — named Reference with a target, named or ISSUER Value; empty Reference, unnamed codecs and other kinds fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestPlausibleEmbedded(t *testing.T) {
	for r, want := range map[*core.Record]bool{
		{Kind: core.KindReference, Codec: core.CodecLocal, Payload: []core.Word{1}}:          true,
		{Kind: core.KindReference, Codec: core.CodecLocal}:                                   false,
		{Kind: core.KindReference, Codec: 0x00, Payload: []core.Word{1}}:                     false,
		{Kind: core.KindValue, Codec: core.CodecUTF8TextEven}:                                true,
		{Kind: core.KindValue, Codec: core.CodecValueIssuer2Word, Issuer: []core.Word{0, 2}}: true,
		{Kind: core.KindValue, Codec: 0x30}:                                                  false,
		{Kind: core.KindEdge, Codec: 0x00}:                                                   false,
	} {
		if plausibleEmbedded(*r) != want {
			t.Fatalf("%+v: want %v", *r, want)
		}
	}
}
