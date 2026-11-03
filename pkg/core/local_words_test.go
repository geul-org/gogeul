//ff:func feature=core type=decoder control=sequence
//ff:what TestLocalWords — the 2-word LOCAL record
package core

import (
	"testing"
)

func TestLocalWords(t *testing.T) {
	if ws := localWords(3); len(ws) != 2 || ws[0] != 0x0051 || ws[1] != 3 {
		t.Fatalf("got %04X", ws)
	}
}
