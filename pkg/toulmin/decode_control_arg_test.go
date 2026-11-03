//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestDecodeControlArg — PROSE span, SECTION code and case REF arguments
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDecodeControlArg(t *testing.T) {
	p, pos := ControlPacket{Code: ControlProse}, 0
	if err := decodeControlArg([]core.Word{1, 2}, &pos, &p, testWire); err != nil || p.Span == nil || p.Span.Len != 2 {
		t.Fatalf("prose: %v %+v", err, p)
	}
	p, pos = ControlPacket{Code: ControlSection}, 0
	if err := decodeControlArg([]core.Word{5}, &pos, &p, testWire); err != nil || p.Arg != 5 {
		t.Fatalf("section: %v %+v", err, p)
	}
	p, pos = ControlPacket{Code: ControlCaseEnd}, 0
	if err := decodeControlArg([]core.Word{0x0051, 7}, &pos, &p, testWire); err != nil || p.Arg != 7 {
		t.Fatalf("case: %v %+v", err, p)
	}
}
