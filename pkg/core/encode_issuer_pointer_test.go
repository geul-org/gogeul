//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestEncodeIssuerPointer — a 2-word Issuer encodes as 0032 + Issuer; 1-word, 3-word and zero-padded Issuers fail
package core

import (
	"slices"
	"testing"
)

func TestEncodeIssuerPointer(t *testing.T) {
	r, err := EncodeIssuerPointer([]Word{0, 0x0D})
	if err != nil {
		t.Fatal(err)
	}
	if ws, _ := EncodeRecord(r); !slices.Equal(ws, []Word{0x0032, 0, 0x0D}) {
		t.Fatalf("%04X", ws)
	}
	for _, bad := range [][]Word{{1, 2, 3}, {0x40}, {0, 0, 0, 1}} {
		if _, err := EncodeIssuerPointer(bad); err == nil {
			t.Fatalf("%04X accepted", bad)
		}
	}
}
