//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestTakeValue — reads an embedded Value and advances; bad records and non-Values fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTakeValue(t *testing.T) {
	pos := 0
	if r, err := TakeValue([]core.Word{0x0421, 7}, &pos); err != nil || r.Codec != 0x02 || pos != 2 {
		t.Fatalf("got %+v %v pos=%d", r, err, pos)
	}
	pos = 0
	if _, err := TakeValue([]core.Word{0xE000}, &pos); err == nil {
		t.Fatal("bad record: want error")
	}
	pos = 0
	if _, err := TakeValue([]core.Word{0x1000}, &pos); err == nil {
		t.Fatal("non-Value: want error")
	}
}
