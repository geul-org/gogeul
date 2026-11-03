//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderHasUnder — a NAMED path answers without keys unless removed; keys answer for Issuers, documents and h; other names have none
package pack

import (
	"context"
	"testing"
)

func TestReaderHasUnder(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	r := a.sets[0].r
	for _, n := range []string{"0x00000000", "0x00000000/docs/0x0000000000000008", "0x00000000/docs", "0x00000000/hashes"} {
		if ok, err := r.hasUnder(ctx, n, nil); err != nil || !ok {
			t.Fatal(n, ok, err)
		}
	}
	for _, n := range []string{"issuers/0x00000001.geul", "zz", "0x00000000/codecs/0x00000420/0x00000001.geul"} {
		if ok, err := r.hasUnder(ctx, n, nil); err != nil || ok {
			t.Fatal(n, ok, err)
		}
	}
	bare := &Reader{m: &manifest{named: []namedEntry{{path: "a/b.geul"}, {path: "c.geul"}}}}
	if ok, _ := bare.hasUnder(ctx, "a", nil); !ok {
		t.Fatal("named")
	}
	if ok, _ := bare.hasUnder(ctx, "a", map[string]bool{"a/b.geul": true}); ok {
		t.Fatal("named, removed")
	}
}
