//ff:func feature=build type=helper control=sequence
//ff:what TestIssuerNames — names of issuers.geul; none when missing or broken
package build

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIssuerNames(t *testing.T) {
	root := namesFixture(t)
	if n := issuerNames(treeOf(root)); len(n) != 1 || n["geul.org"] != "0x00000000" {
		t.Fatalf("got %v", n)
	}
	os.WriteFile(filepath.Join(root, ".well-known", "geul", "issuers/0x00000001.geul"), []byte{1}, 0o644)
	if n := issuerNames(treeOf(root)); len(n) != 0 {
		t.Fatalf("broken: %v", n)
	}
	if n := issuerNames(treeOf(t.TempDir())); len(n) != 0 {
		t.Fatalf("missing: %v", n)
	}
}
