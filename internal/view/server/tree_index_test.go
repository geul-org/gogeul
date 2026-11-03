//ff:func feature=server type=helper control=sequence
//ff:what TestTreeIndex — a tree directory listed with its Issuer named; a missing directory fails
package server

import "testing"

func TestTreeIndex(t *testing.T) {
	root := namesFixture(t)
	p, err := treeIndex(treeOf(root), "/.well-known/geul/")
	if err != nil || p.Path != "/.well-known/geul/" || len(p.Entries) != 2 || p.Entries[0] != (indexEntry{Name: "geul.org/", URL: "/.well-known/geul/0x00000000/", ID: "0x00000000"}) {
		t.Fatalf("top: %+v %v", p, err)
	}
	if p, _ := treeIndex(treeOf(root), "/.well-known/geul/0x00000000/"); p.Entries[1].Name != "docs/" {
		t.Fatalf("issuer: %+v", p)
	}
	if p, _ := treeIndex(treeOf(root), "/.well-known/geul/0x00000000/codecs/"); len(p.Entries) != 1 || p.Entries[0].Name != "0x00000000/" || p.Entries[0].ID != "" {
		t.Fatalf("codec ID not named as an Issuer: %+v", p)
	}
	if _, err := treeIndex(treeOf(root), "/.well-known/geul/none/"); err == nil {
		t.Fatal("missing: want error")
	}
}
