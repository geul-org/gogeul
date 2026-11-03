//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseIssuers — bad count, a 1-word or other width, truncated ID, a zero-padded entry and duplicates fail
package pack

import (
	"errors"
	"testing"
)

func TestParseIssuers(t *testing.T) {
	bad := [][]byte{
		{9, 0, 0, 0},
		{1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 2, 0, 0},
		{1, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 1},
		{2, 0, 0, 0, 2, 0, 0, 0, 1, 2, 0, 0, 0, 1},
		{1},
	}
	for i, b := range bad {
		if _, err := parseIssuers(&cursor{b: b}); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
	if got, err := parseIssuers(&cursor{b: []byte{1, 0, 0, 0, 4, 0, 1, 0, 0, 0, 0, 0, 0}}); err != nil || got[0] != "0x0001000000000000" {
		t.Fatal(got, err)
	}
}
