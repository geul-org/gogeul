//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeWith — UNSIGNED-INTEGER count then one LOCAL REF per item; empty list is the count alone
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEncodeWith(t *testing.T) {
	want := []core.Word{0x0421, 2, 0x0051, 5, 0x0051, 6}
	if got := encodeWith(nil, []TID{5, 6}); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %04X", got)
	}
	if got := encodeWith([]core.Word{9}, nil); !reflect.DeepEqual(got, []core.Word{9, 0x0421, 0}) {
		t.Fatalf("empty: %04X", got)
	}
}
