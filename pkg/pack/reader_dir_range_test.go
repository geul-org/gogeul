//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderDirRange — `h`, an Issuer and a document by shape; the last document ends at the next Issuer; unknown Issuers and other names have no keys
package pack

import "testing"

func TestReaderDirRange(t *testing.T) {
	r := &Reader{m: &manifest{}, inum: map[string]uint32{"0x0000": 3}}
	if kr, ok := r.dirRange("h"); !ok || kr.lo != "" || kr.hi != "" || kr.path("\x01") != hPath("\x01") {
		t.Fatalf("h: %+v", kr)
	}
	kr, ok := r.dirRange("0x0000")
	if !ok || kr.lo != pKey(3, 0, 0) || kr.hi != pKey(4, 0, 0) {
		t.Fatalf("Issuer: %+v", kr)
	}
	if got := kr.path(pKey(3, 8, 0)); got != "0x0000/0x0000000000000008/versions.geul" {
		t.Fatal(got)
	}
	kr, ok = r.dirRange("0x0000/0x0000000000000008")
	if !ok || kr.lo != pKey(3, 8, 0) || kr.hi != pKey(3, 9, 0) {
		t.Fatalf("document: %+v", kr)
	}
	kr, ok = r.dirRange("0x0000/0xffffffffffffffff")
	if !ok || kr.lo != pKey(3, ^uint64(0), 0) || kr.hi != pKey(4, 0, 0) {
		t.Fatalf("last document: %+v", kr)
	}
	for _, name := range []string{"0x0001", "0x0001/0x0000000000000008", "0x0000/zz", "0x0000/0x000000000000000A", "0x0000/0x0000000000000008/versions.geul", "issuers.geul"} {
		if _, ok := r.dirRange(name); ok {
			t.Fatal(name)
		}
	}
}
