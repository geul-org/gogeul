//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestHeadIssuer — the Issuer of a document DOC-ID or a codec DOC-ID; other HEAD records have none
package pack

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHeadIssuer(t *testing.T) {
	doc, _ := core.EncodeDocID([]core.Word{0, 2}, []core.Word{0, 0, 0, 1})
	codec, _ := core.EncodeDocIDCodec([]core.Word{0, 3}, 0x0B80)
	for want, r := range map[core.Word]core.Record{2: doc, 3: codec} {
		if is, ok := headIssuer(r); !ok || !slices.Equal(is, []core.Word{0, want}) {
			t.Fatalf("%v: %v %v", r, is, ok)
		}
	}
	if _, ok := headIssuer(core.EncodeUintValue(1)); ok {
		t.Fatal("Value: want none")
	}
}
