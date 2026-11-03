//ff:func feature=server type=helper control=iteration dimension=1
//ff:what TestTitleOf — Issuer, document and codec spec IDs shown by name under .well-known/geul/; unnamed paths unchanged
package server

import (
	"testing"
)

func TestTitleOf(t *testing.T) {
	root := namesFixture(t)
	for p, want := range map[string]string{
		"/.well-known/geul/0x00000000/docs/0xffff000000000002/0x00000001.geul": ".well-known/geul/geul.org/sample-b/0x00000001.geul",
		"/.well-known/geul/0x00000000/docs/0x00000001.geul":                    ".well-known/geul/geul.org/docs/0x00000001.geul",
		"/.well-known/geul/h/0xab.geul":                                        ".well-known/geul/h/0xab.geul",
		"/.well-known/geul/0x00000000/codecs/0x00000980/0x00000001.geul":       ".well-known/geul/geul.org/publication/0x00000001.geul",
	} {
		if got := titleOf(treeOf(root), p); got != want {
			t.Fatalf("%s: %q", p, got)
		}
	}
}
