//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveCompact — each set gets a fresh generation serving the same paths; staged changes, remote archives and failures stop it
package pack

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestArchiveCompact(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	a.set("0x00000002")
	before, _ := a.Paths()
	rs, err := a.Compact(ctx)
	after, _ := a.Paths()
	if err != nil || len(rs) != 1 || a.sets[0].r.Generation() != 2 || !reflect.DeepEqual(before, after) {
		t.Fatal(rs, err)
	}
	a.WriteFile("x", nil)
	if _, err := a.Compact(ctx); err == nil {
		t.Fatal("staged")
	}
	if _, err := (&Archive{}).Compact(ctx); err == nil {
		t.Fatal("remote")
	}
	b, _ := testArchive(t)
	b.o.Codec = fakeCodec{err: errors.New("x")}
	if _, err := b.Compact(ctx); err == nil {
		t.Fatal("write")
	}
}
