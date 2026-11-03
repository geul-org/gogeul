//ff:func feature=inspect type=decoder control=sequence
//ff:what TestParseDocID — the Issuer and document of the Reference; a raw payload fails
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseDocID(t *testing.T) {
	r, _ := core.EncodeDocID([]core.Word{0, 2}, []core.Word{0, 0, 0, 2})
	d, ok := parseDocID(r)
	if !ok || len(d.Issuer) != 2 || d.Issuer[1] != 2 || len(d.Doc) != 4 || d.Doc[3] != 2 {
		t.Fatalf("got %+v %v", d, ok)
	}
	c, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, 0x0980)
	if d, ok := parseDocID(c); !ok || len(d.Doc) != 2 || d.Doc[1] != 0x0980 {
		t.Fatalf("codec: %+v %v", d, ok)
	}
	if _, ok := parseDocID(core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: []core.Word{1, 0, 0, 0, 2}}); ok {
		t.Fatal("raw payload: want not ok")
	}
}
