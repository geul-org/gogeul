//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestParsePattern — Issuer, Issuer/doc, Issuer/*, *; bad Issuer and bad document fail
package resolve

import (
	"testing"
)

func TestParsePattern(t *testing.T) {
	cases := map[string]Pattern{
		"0x00001234":                    {Issuer: "0x00001234"},
		"0x00001234/0x00000000000000ab": {Issuer: "0x00001234", Doc: "0x00000000000000ab"},
		"0x00001234/*":                  {Issuer: "0x00001234"},
		"*":                             {Issuer: "*"},
	}
	for s, want := range cases {
		if p, err := parsePattern(s); err != nil || p != want {
			t.Fatalf("%q: %+v %v", s, p, err)
		}
	}
	for _, s := range []string{"0x123", "0x0001", "0x00001234/0xab", "0x00001234/0x00000000000000AB"} {
		if _, err := parsePattern(s); err == nil {
			t.Fatalf("%q: want error", s)
		}
	}
}
