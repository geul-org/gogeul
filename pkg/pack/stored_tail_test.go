//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestStoredTail — versions.geul and a nonzero 8-digit version file are stored names; latest.geul, version 0 and other spellings are not
package pack

import "testing"

func TestStoredTail(t *testing.T) {
	for n, want := range map[string]bool{
		"versions.geul": true, "0x00000001.geul": true, "0xFFFFFFFF.geul": true,
		"latest.geul": false, "0x00000000.geul": false, "0x00000001": false, "0x0001.geul": false, "0x0000000A.geul": true, "0x0000000a.geul": true, "0X0000000A.geul": false,
	} {
		if storedTail(n) != want {
			t.Fatalf("%s", n)
		}
	}
}
