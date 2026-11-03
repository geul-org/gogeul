//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderDirRange — <I> and <I>/docs are the Issuer's P keys, <I>/docs/<D> one document's, <I>/hashes the H keys; other names have none
package pack

import "testing"

func TestReaderDirRange(t *testing.T) {
	r := &Reader{m: &manifest{}, inum: map[string]uint32{"0x00000000": 3}}
	if kr, ok := r.dirRange("0x00000000/hashes"); !ok || kr.lo != "" || kr.hi != "" || kr.path("\x01") != hPath("0x00000000", "\x01") {
		t.Fatalf("hashes: %+v", kr)
	}
	for _, name := range []string{"0x00000000", "0x00000000/docs", "0x0000000000000000/docs"} {
		kr, ok := r.dirRange(name)
		if !ok || kr.lo != pKey(3, 0, 0) || kr.hi != pKey(4, 0, 0) {
			t.Fatalf("%s: %+v", name, kr)
		}
		if got := kr.path(pKey(3, 8, 0)); got != "0x00000000/docs/0x0000000000000008/versions.geul" {
			t.Fatal(got)
		}
	}
	kr, ok := r.dirRange("0x00000000/docs/0x0000000000000008")
	if !ok || kr.lo != pKey(3, 8, 0) || kr.hi != pKey(3, 9, 0) {
		t.Fatalf("document: %+v", kr)
	}
	kr, ok = r.dirRange("0x00000000/docs/0xFFFFFFFFFFFFFFFF")
	if !ok || kr.lo != pKey(3, ^uint64(0), 0) || kr.hi != pKey(4, 0, 0) {
		t.Fatalf("last document: %+v", kr)
	}
	kr, ok = r.dirRange("0x00000000/docs/0x000000000000000a")
	if !ok || kr.lo != pKey(3, 10, 0) || kr.hi != pKey(3, 11, 0) {
		t.Fatalf("lower case document: %+v", kr)
	}
	for _, name := range []string{"0x00000001", "0x00000001/docs/0x0000000000000008", "0x00000000/zz", "0x00000000/docs/0X000000000000000A",
		"0x00000000/docs/0x0000000000000008/versions.geul", "issuers", "0x00000000/codecs", "h"} {
		if _, ok := r.dirRange(name); ok {
			t.Fatal(name)
		}
	}
}
