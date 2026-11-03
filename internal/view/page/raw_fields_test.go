//ff:func feature=page type=helper control=sequence
//ff:what TestRawFields — an ISSUER record shows Issuer and format; others show nothing
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRawFields(t *testing.T) {
	f := rawFields(core.Record{Issuer: []core.Word{0, 0}, Format: 0xC0FF})
	if len(f) != 2 || f[0].Text != "0x00000000" || f[1].Text != "0x0000C0FF" {
		t.Fatalf("got %+v", f)
	}
	if f := rawFields(core.Record{Codec: 0x61}); f != nil {
		t.Fatalf("direct: %+v", f)
	}
}
