//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestNormalIssuer — the issued-width spelling of 8/16/64-digit Issuers; other lengths and non-hex fail
package resolve

import "testing"

func TestNormalIssuer(t *testing.T) {
	for in, want := range map[string]string{"0x0000000000000005": "0x00000005", "0x0000000100000000": "0x0000000100000000", "0x00000005": "0x00000005"} {
		if got, ok := normalIssuer(in); !ok || got != want {
			t.Fatal(in, got, ok)
		}
	}
	for _, in := range []string{"0x0005", "0x00000000000000000005", "0xzzzzzzzz"} {
		if _, ok := normalIssuer(in); ok {
			t.Fatal(in)
		}
	}
}
