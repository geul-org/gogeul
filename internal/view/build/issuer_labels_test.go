//ff:func feature=build type=helper control=sequence
//ff:what TestIssuerLabels — the inverted names of issuers.geul; none when the tree has no issuers.geul
package build

import "testing"

func TestIssuerLabels(t *testing.T) {
	if l := issuerLabels(treeOf(namesFixture(t))); len(l) != 1 || l["0x00000000"] != "geul.org" {
		t.Fatalf("got %v", l)
	}
	if l := issuerLabels(treeOf(t.TempDir())); len(l) != 0 {
		t.Fatalf("none: %v", l)
	}
}
