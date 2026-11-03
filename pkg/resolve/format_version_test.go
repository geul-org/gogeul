//ff:func feature=resolve type=encoder control=sequence
//ff:what TestFormatVersion — 8 upper-case hex digits
package resolve

import (
	"testing"
)

func TestFormatVersion(t *testing.T) {
	if s := FormatVersion(0xAB); s != "0x000000AB" {
		t.Fatalf("got %q", s)
	}
}
