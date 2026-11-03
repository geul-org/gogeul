//ff:func feature=build type=helper control=iteration dimension=1
//ff:what TestTreeLabels — Issuer names at the top of the tree (also for its index.json path); nothing below, so a codec ID that looks like an Issuer ID keeps its ID
package build

import (
	"testing"
)

func TestTreeLabels(t *testing.T) {
	root := namesFixture(t)
	for _, p := range []string{"/.well-known/geul/", "/.well-known/geul"} {
		if l := treeLabels(treeOf(root), p); len(l) != 1 || l["0x00000000"] != "geul.org" {
			t.Fatalf("top %q: %v", p, l)
		}
	}
	for _, p := range []string{"/.well-known/geul/0x00000000/", "/.well-known/geul/0x00000000/codecs"} {
		if l := treeLabels(treeOf(root), p); len(l) != 0 {
			t.Fatalf("%s: %v", p, l)
		}
	}
}
