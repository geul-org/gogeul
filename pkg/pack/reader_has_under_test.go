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
	for _, n := range []string{"0x0000", "0x0000/0x0000000000000008", "h"} {
		if ok, err := r.hasUnder(ctx, n, nil); err != nil || !ok {
			t.Fatal(n, ok, err)
		}
	}
	for _, n := range []string{"issuers.geul", "zz", "0x0000/documents.geul"} {
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
