//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseMeta — a named Meta with tokens, a whole raw Meta; no name, an unknown name, a raw non-Meta and bad tokens fail
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseMeta(t *testing.T) {
	r, err := ParseMeta(`message "first version"`, Context{})
	if err != nil || r.Kind != core.KindMeta || r.Codec != core.MetaMessage {
		t.Fatalf("%+v %v", r, err)
	}
	sv, _ := core.EncodeRecord(core.EncodeSpecVersion(2))
	if r, err := ParseMeta(hexToken(sv), Context{}); err != nil || r.Codec != core.MetaSpecVersion {
		t.Fatalf("raw: %+v %v", r, err)
	}
	for _, s := range []string{"", "NOPE 1", "0x00510001", "MESSAGE word", `MESSAGE "open`} {
		if _, err := ParseMeta(s, Context{}); err == nil {
			t.Errorf("%q: want error", s)
		}
	}
}
