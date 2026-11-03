//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveCommitSet — writes and reopens; an unchanged batch keeps the reader; read-only archives and write errors fail
package pack

import (
	"context"
	"testing"
)

func TestArchiveCommitSet(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	s := a.sets[0]
	old := s.r
	get := memGet(map[string][]byte{"issuers/0x00000002.geul": []byte("x"), "issuers/0x00000001.geul": fixtureFile(t, "issuers/0x00000001.geul")})
	if res, err := a.commitSet(ctx, s, &Batch{Put: []string{"issuers/0x00000001.geul"}, Get: get}); err != nil || res.Changed || s.r != old {
		t.Fatal(res, err)
	}
	if res, err := a.commitSet(ctx, s, &Batch{Put: []string{"issuers/0x00000002.geul"}, Get: get}); err != nil || s.r.Generation() != 2 {
		t.Fatal(res, err)
	}
	if _, err := a.commitSet(ctx, s, &Batch{Put: []string{"issuers/0x00000003.geul"}, Get: get}); err == nil {
		t.Fatal("get")
	}
	if _, err := (&Archive{}).commitSet(ctx, s, &Batch{}); err == nil {
		t.Fatal("read-only")
	}
}
