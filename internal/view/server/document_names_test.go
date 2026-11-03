//ff:func feature=server type=helper control=sequence
//ff:what TestDocumentNames — names of that Issuer's document list; other Issuers and missing files give none
package server

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDocumentNames(t *testing.T) {
	root := namesFixture(t)
	if n := documentNames(treeOf(root), "0x00000000"); len(n) != 2 || n["sample-b"] != "0xffff000000000002" || n["publication"] != "0x00000980" {
		t.Fatalf("got %v", n)
	}
	other := filepath.Join(root, ".well-known", "geul", "0x00000002")
	os.MkdirAll(filepath.Join(other, "docs"), 0o755)
	b, _ := os.ReadFile(filepath.Join(root, ".well-known", "geul", "0x00000000", "docs", "0x00000001.geul"))
	os.WriteFile(filepath.Join(other, "docs", "0x00000001.geul"), b, 0o644)
	if n := documentNames(treeOf(root), "0x00000002"); len(n) != 0 {
		t.Fatalf("foreign: %v", n)
	}
	if n := documentNames(treeOf(root), "0x00000003"); len(n) != 0 {
		t.Fatalf("missing: %v", n)
	}
}
