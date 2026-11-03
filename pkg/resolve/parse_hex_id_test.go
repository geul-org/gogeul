//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestParseHexID — canonical IDs decode; either case decodes; wrong width, prefix or digits fail
package resolve

import (
	"testing"
)

func TestParseHexID(t *testing.T) {
	if b, ok := parseHexID("0x00AB", 1); !ok || len(b) != 2 || b[1] != 0xab {
		t.Fatalf("good: %x %v", b, ok)
	}
	for _, s := range []string{"0x00ab", "0x00Ab"} {
		if b, ok := parseHexID(s, 1); !ok || b[1] != 0xab {
			t.Fatalf("%q either case: %x %v", s, b, ok)
		}
	}
	for _, s := range []string{"0x0AB", "0X00ab", "0X00AB", "0x00zz", "1x00ab"} {
		if _, ok := parseHexID(s, 1); ok {
			t.Fatalf("%q: want not ok", s)
		}
	}
}
