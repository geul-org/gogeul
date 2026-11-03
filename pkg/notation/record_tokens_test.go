//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestRecordTokens — records with short forms and addresses split; a broken record, a raw-only one, a slot marker and nothing are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordTokens(t *testing.T) {
	doc, _ := core.EncodeDocPointer([]core.Word{0, 0}, []core.Word{0, 0, 0, 1}, []core.Word{0, 0})
	ws, _ := core.EncodeRecord(doc)
	if got, ok := recordTokens(ws); !ok || got[0] != "0x00000000/0x0000000000000001" {
		t.Fatal(got)
	}
	for _, x := range [][]core.Word{{0x0B86}, {0x0002}, {0x0051, 3, 0x0001}, nil} {
		if _, ok := recordTokens(x); ok {
			t.Errorf("%04X: want false", x)
		}
	}
}
