//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestEncodeControl — SECTION golden plus retired-control, PROSE and section-code errors
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestEncodeControl pins a SECTION record (EDGE-ISSUER-2WORD, GEUL Issuer 0x00000000,
// SECTION direct code 0x102L (§3.14), Edge Id, section code) and checks every error.
func TestEncodeControl(t *testing.T) {
	got, err := EncodeControl(ControlPacket{ID: 3, Code: ControlSection, Arg: 5}, testWire)
	if want := []core.Word{0x1023, 0x0051, 3, 5}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("section: %04X %v want %04X", got, err, want)
	}
	bad := []ControlPacket{
		{Code: 0},
		{Code: ControlProse},
		{Code: ControlProse, Span: &Span{Off: 0x10000}},
		{Code: ControlSection, Arg: 0x10000},
	}
	for i, p := range bad {
		if _, err := EncodeControl(p, testWire); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
