//ff:func feature=schema type=rule control=sequence
//ff:what TestNoLinguisticContent — zxx in any case, with subtags or odd padding, is no linguistic content; und, mul, en and empty are not
package schema

import "testing"

func TestNoLinguisticContent(t *testing.T) {
	if !noLinguisticContent([]byte("zxx\x00"), true) || !noLinguisticContent([]byte("ZXX-Latn"), false) {
		t.Fatal("zxx")
	}
	if noLinguisticContent([]byte("und\x00"), true) || noLinguisticContent([]byte("mul\x00"), true) || noLinguisticContent([]byte("en"), false) || noLinguisticContent(nil, true) {
		t.Fatal("not zxx")
	}
}
