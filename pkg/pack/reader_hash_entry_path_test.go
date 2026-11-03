//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderHashEntryPath — a sole-Issuer set names a hash under that Issuer without reading; a several-Issuer set reads the record (HashPath) and fails when it cannot
package pack

import (
	"context"
	"strings"
	"testing"
)

func TestReaderHashEntryPath(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	k := strings.Repeat("\xab", 32)
	if p, err := r.hashEntryPath(ctx, k, Loc{Obj: 99}); err != nil || p != hPath("0x00000000", k) {
		t.Fatal(p, err)
	}
	l, _ := r.find(ctx, r.m.pLayers, pKey(0, 0x845, 1), "")
	r.m.issuers = append(r.m.issuers, "0x00000002")
	h, _ := contentHash(fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul"))
	if p, err := r.hashEntryPath(ctx, k, l); err != nil || p != hPath("0x00000000", string(h[:])) {
		t.Fatal(p, err)
	}
	if _, err := r.hashEntryPath(ctx, k, Loc{Obj: 99}); err == nil {
		t.Fatal("unreadable: want error")
	}
}
