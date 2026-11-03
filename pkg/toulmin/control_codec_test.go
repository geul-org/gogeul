//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestControlCodec — every control round-trips through ControlFromCodec; unknown fails both ways
package toulmin

import "testing"

func TestControlCodec(t *testing.T) {
	for _, c := range controls {
		f, err := ControlCodec(c)
		if err != nil {
			t.Fatal(err)
		}
		if got, ok := ControlFromCodec(f); !ok || got != c {
			t.Fatalf("0x%02X round trip got 0x%02X", uint16(c), uint16(got))
		}
	}
	if _, err := ControlCodec(0); err == nil {
		t.Fatal("retired control must fail")
	}
	if _, ok := ControlFromCodec(CodecText); ok {
		t.Fatal("TEXT is not a control")
	}
}
