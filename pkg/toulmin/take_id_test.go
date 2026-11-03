//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestTakeID — reads a LOCAL Reference and advances; truncation and a non-LOCAL record fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTakeID(t *testing.T) {
	pos := 1
	if v, err := TakeID([]core.Word{9, 0x0052, 1, 2}, &pos); err != nil || v != 0x10002 || pos != 4 {
		t.Fatalf("got %d %v pos=%d", v, err, pos)
	}
	pos = 0
	if _, err := TakeID([]core.Word{0x0052, 1}, &pos); err == nil {
		t.Fatal("truncated: want error")
	}
	pos = 0
	if _, err := TakeID([]core.Word{0x0621, 1}, &pos); err == nil || pos != 0 {
		t.Fatal("Value record: want error")
	}
}
