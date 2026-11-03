//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestDecodeWith — a counted REF list decodes; missing or wrong count, oversized count and a bad REF fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDecodeWith(t *testing.T) {
	pos := 0
	if got, err := decodeWith([]core.Word{0x0421, 2, 0x0051, 5, 0x0051, 6}, &pos); err != nil || len(got) != 2 || got[1] != 6 || pos != 6 {
		t.Fatalf("got %v %v", got, err)
	}
	pos = 0
	if _, err := decodeWith(nil, &pos); err == nil {
		t.Fatal("missing count: want error")
	}
	pos = 0
	if _, err := decodeWith([]core.Word{0x0890, 0, 2}, &pos); err == nil {
		t.Fatal("non-integer count: want error")
	}
	pos = 0
	if _, err := decodeWith([]core.Word{0x0421, 3, 0x0051, 5}, &pos); err == nil {
		t.Fatal("count past payload: want error")
	}
	pos = 0
	if _, err := decodeWith([]core.Word{0x0421, 1, 0x0421, 5}, &pos); err == nil {
		t.Fatal("non-LOCAL REF: want error")
	}
}
