//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestStoredHexID — a stored folder name is an ID only in A-F upper case; lower case, other widths and `0X` are not
package resolve

import "testing"

func TestStoredHexID(t *testing.T) {
	if b, ok := storedHexID("0x00000ABC", 2); !ok || b[2] != 0x0A || b[3] != 0xBC {
		t.Fatalf("upper: %x %v", b, ok)
	}
	for _, s := range []string{"0x00000abc", "0x00000aBC", "0X00000ABC", "0x0ABC", "0x00000ABG"} {
		if _, ok := storedHexID(s, 2); ok {
			t.Fatalf("%q: want not ok", s)
		}
	}
}
