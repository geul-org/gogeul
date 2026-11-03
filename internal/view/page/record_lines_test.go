//ff:func feature=page type=helper control=sequence
//ff:what TestRecordLines — a Value keeps one payload run; a Reference names its fields; a long Reference shows its Length word
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordLines(t *testing.T) {
	if l := recordLines([]core.Word{0x0621, 5}); len(l) != 2 || l[0].Name != "TAG" || l[1].Name != "payload" {
		t.Fatalf("value: %+v", l)
	}
	if l := recordLines([]core.Word{0x0051, 7}); l[1].Name != "대상 ID" {
		t.Fatalf("local: %+v", l)
	}
	long := append([]core.Word{0x005E, 16}, make([]core.Word, 14)...)
	long[2] = 1
	if l := recordLines(long); len(l) != 16 || l[1].Name != "LENGTH" || l[2].Name != "대상 ID" {
		t.Fatalf("long: %+v", l[:3])
	}
}
