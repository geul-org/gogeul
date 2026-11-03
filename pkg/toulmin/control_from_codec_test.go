//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestControlFromCodec — known control codecs resolve; other codecs do not
package toulmin

import "testing"

func TestControlFromCodec(t *testing.T) {
	if c, ok := ControlFromCodec(CodecProse); !ok || c != ControlProse {
		t.Fatalf("PROSE: %v %v", c, ok)
	}
	if _, ok := ControlFromCodec(CodecText); ok {
		t.Fatal("TEXT is not a control")
	}
}
