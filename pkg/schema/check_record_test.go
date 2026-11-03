//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckRecord — dispatch by Kind
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckRecord(t *testing.T) {
	cases := map[string]core.Record{
		"R1": {Kind: core.KindReference, Codec: 0x00},
		"V2": {Kind: core.KindValue, Codec: core.CodecBool, Payload: []core.Word{2}},
		"E1": {Kind: core.KindEdge, Codec: 0x00},
		"M4": {Kind: core.KindMeta, Codec: core.MetaSpecVersion},
	}
	for want, r := range cases {
		if got := rulesOf(checkRecord(recOf(r), nil)); got != want {
			t.Fatalf("%s: %q", want, got)
		}
	}
}
