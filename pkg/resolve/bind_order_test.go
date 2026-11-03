//ff:func feature=resolve type=resolver control=sequence
//ff:what TestBindOrder — names per document in record order; other records and non-document BINDs are skipped
package resolve

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestBindOrder(t *testing.T) {
	doc, _ := core.EncodeDocPointer([]core.Word{0, 0}, []core.Word{0, 0, 0, 7}, []core.Word{0, 0})
	is, _ := core.EncodeIssuerPointer([]core.Word{0, 0})
	b1, _ := core.EncodeBind(core.Bind{EdgeID: 1, Subject: doc, Name: []byte("b")})
	b2, _ := core.EncodeBind(core.Bind{EdgeID: 2, Subject: is, Name: []byte("i")})
	b3, _ := core.EncodeBind(core.Bind{EdgeID: 3, Subject: doc, Name: []byte("a")})
	got := bindOrder([]core.Record{b1, core.EncodeSpecVersion(2), b2, b3})
	if len(got) != 1 || !slices.Equal(got["0x0000000000000007"], []string{"b", "a"}) {
		t.Fatal(got)
	}
}
