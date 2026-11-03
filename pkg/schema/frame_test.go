//ff:func feature=schema type=rule control=sequence
//ff:what TestFrame — records in order; framing stops at an untrusted Length
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestFrame(t *testing.T) {
	recs, fs := frame([]core.Word{0x0421, 5, 0x0421, 6})
	if len(recs) != 2 || fs != nil || recs[1].Word != 2 || recs[1].Index != 1 {
		t.Fatalf("ok: %+v %v", recs, fs)
	}
	if recs, fs := frame([]core.Word{0x0421, 5, 0x040E}); len(recs) != 1 || rulesOf(fs) != "F2" {
		t.Fatalf("stop: %+v %v", recs, fs)
	}
}
