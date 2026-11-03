//ff:func feature=resolve type=decoder control=sequence
//ff:what TestParsePatterns — comma list with blanks; one bad item fails
package resolve

import (
	"testing"
)

func TestParsePatterns(t *testing.T) {
	ps, err := parsePatterns(" 0x00000002 ,, * ")
	if err != nil || len(ps) != 2 || ps[1].Issuer != "*" {
		t.Fatalf("got %+v %v", ps, err)
	}
	if _, err := parsePatterns("0x00000002,bad"); err == nil {
		t.Fatal("bad: want error")
	}
}
