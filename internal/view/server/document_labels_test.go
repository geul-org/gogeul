//ff:func feature=server type=helper control=sequence
//ff:what TestDocumentLabels — document ID → name of the Issuer's own documents; none when the document list is missing
package server

import "testing"

func TestDocumentLabels(t *testing.T) {
	root := namesFixture(t)
	if l := documentLabels(treeOf(root), "0x00000000"); l["0xffff000000000002"] != "sample-b" || len(l) != 2 {
		t.Fatal(l)
	}
	if l := documentLabels(treeOf(root), "0x0002"); len(l) != 0 {
		t.Fatal(l)
	}
}
