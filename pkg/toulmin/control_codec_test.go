//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestControlFormat — every control round-trips through ControlFromFormat; unknown fails both ways
package toulmin

import "testing"

func TestControlFormat(t *testing.T) {
	for _, c := range controls {
		f, err := ControlFormat(c)
		if err != nil {
			t.Fatal(err)
		}
		if got, ok := ControlFromFormat(f); !ok || got != c {
			t.Fatalf("0x%02X round trip got 0x%02X", uint16(c), uint16(got))
		}
	}
	if _, err := ControlFormat(0); err == nil {
		t.Fatal("retired control must fail")
	}
	if _, ok := ControlFromFormat(FormatText); ok {
		t.Fatal("TEXT is not a control")
	}
}
