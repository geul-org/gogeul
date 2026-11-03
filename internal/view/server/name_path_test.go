//ff:func feature=server type=helper control=sequence
//ff:what TestNamePath — an Issuer name becomes its ID and the rest is kept; a document name stays as it is; a path with no Issuer name is false
package server

import (
	"testing"
)

func TestNamePath(t *testing.T) {
	root := namesFixture(t)
	if p, ok := namePath(treeOf(root), "/.well-known/geul/geul.org/docs/0xffff000000000002/x"); !ok || p != "/.well-known/geul/0x00000000/docs/0xffff000000000002/x" {
		t.Fatalf("issuer and path: %q %v", p, ok)
	}
	if p, ok := namePath(treeOf(root), "/.well-known/geul/geul.org"); !ok || p != "/.well-known/geul/0x00000000" {
		t.Fatalf("issuer: %q %v", p, ok)
	}
	if p, ok := namePath(treeOf(root), "/.well-known/geul/geul.org/sample-b/"); !ok || p != "/.well-known/geul/0x00000000/sample-b/" {
		t.Fatalf("no document names: %q %v", p, ok)
	}
	if _, ok := namePath(treeOf(root), "/.well-known/geul/0x00000000/sample-b"); ok {
		t.Fatal("no Issuer name: want false")
	}
}
