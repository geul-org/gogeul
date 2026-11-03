//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderHashOwner — a set's sole Issuer; else the document's DOC-ID or version list Issuer; else the first Issuer; 0x00000000 for an empty table
package pack

import "testing"

func TestReaderHashOwner(t *testing.T) {
	r := &Reader{m: &manifest{issuers: []string{"0x00000005"}}}
	doc := fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul")
	if got := r.hashOwner(doc); got != "0x00000005" {
		t.Fatal("sole Issuer", got)
	}
	r.m.issuers = []string{"0x00000005", "0x00000000"}
	if got := r.hashOwner(doc); got != "0x00000000" {
		t.Fatal("DOC-ID", got)
	}
	if got := r.hashOwner([]byte{1}); got != "0x00000005" {
		t.Fatal("unreadable: first Issuer", got)
	}
	if got := r.hashOwner(nil); got != "0x00000005" {
		t.Fatal("no bytes: first Issuer", got)
	}
	r.m.issuers = nil
	if got := r.hashOwner([]byte{1}); got != "0x00000000" {
		t.Fatal("empty table", got)
	}
}
