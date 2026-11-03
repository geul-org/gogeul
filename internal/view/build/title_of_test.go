//ff:func feature=build type=helper control=iteration dimension=1
//ff:what TestTitleOf — the Issuer ID shown by name under .well-known/geul/; document and codec IDs and unnamed paths unchanged
package build

import (
	"testing"
)

func TestTitleOf(t *testing.T) {
	root := namesFixture(t)
	for p, want := range map[string]string{
		"/.well-known/geul/0x00000000/docs/0xFFFF000000000002/0x00000001.geul": ".well-known/geul/geul.org/docs/0xFFFF000000000002/0x00000001.geul",
		"/.well-known/geul/h/0xAB.geul":                                        ".well-known/geul/h/0xAB.geul",
		"/.well-known/geul/0x00000000/codecs/0x00000000/0x00000001.geul":       ".well-known/geul/geul.org/codecs/0x00000000/0x00000001.geul",
	} {
		if got := titleOf(treeOf(root), p); got != want {
			t.Fatalf("%s: %q", p, got)
		}
	}
}
