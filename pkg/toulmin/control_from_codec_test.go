//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestControlFromFormat — known control formats resolve; other formats do not
package toulmin

import "testing"

func TestControlFromFormat(t *testing.T) {
	if c, ok := ControlFromFormat(FormatProse); !ok || c != ControlProse {
		t.Fatalf("PROSE: %v %v", c, ok)
	}
	if _, ok := ControlFromFormat(FormatText); ok {
		t.Fatal("TEXT is not a control")
	}
}
