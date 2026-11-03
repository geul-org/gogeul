//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestFetchAsOf — the version in effect at each instant; before the first, after a sunset, deleted versions skipped, and broken sources fail
package resolve

import (
	"context"
	"errors"
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestFetchAsOf(t *testing.T) {
	ctx := context.Background()
	e := Entry{Name: "law", Source: sampleLawTree(t)}
	k := sampleKey(5, 0)
	for at, want := range map[int64]uint32{150: 1, 250: 2, 450: 3} {
		if n, err := fetchAsOf(ctx, e, k, day(at)); err != nil || n.Version != want {
			t.Fatalf("day %d: %+v %v", at, n, err)
		}
	}
	for _, at := range []int64{50, 350} {
		if _, err := fetchAsOf(ctx, e, k, day(at)); !errors.Is(err, ErrNotFound) {
			t.Fatalf("day %d: %v", at, err)
		}
	}
	if _, err := fetchAsOf(ctx, Entry{Source: Memory{}}, k, day(150)); !errors.Is(err, ErrNotFound) {
		t.Fatalf("no list: %v", err)
	}
	m := Memory{listPath(k.DocKey()): listDocOf(5, VersionEntry{Version: 1})}
	if _, err := fetchAsOf(ctx, Entry{Source: m}, k, day(150)); !errors.Is(err, ErrNotFound) {
		t.Fatalf("no file: %v", err)
	}
	gone := Memory{listPath(k.DocKey()): listDocOf(5, VersionEntry{Version: 1, Deleted: true}), versionPath(k.DocKey(), 1): []byte("junk")}
	if _, err := fetchAsOf(ctx, Entry{Source: gone}, k, day(150)); !errors.Is(err, ErrNotFound) {
		t.Fatalf("deleted not skipped: %v", err)
	}
	bad := Dir{Root: t.TempDir()}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	in := enc(func() core.Record {
		r, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, 5})
		return r
	}())
	in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaEffectiveFrom, Payload: enc(core.EncodeUintValue(1))})...)
	if _, err := bad.Publish(core.PutWords(core.WrapDocument(append(in, 0x1C30, 0x1C40)))); err != nil {
		t.Fatal(err)
	}
	if _, err := fetchAsOf(ctx, Entry{Source: bad}, k, day(150)); !errors.Is(err, ErrVerify) {
		t.Fatalf("bad bound: %v", err)
	}
}
