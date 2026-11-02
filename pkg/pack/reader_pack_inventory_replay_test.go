//ff:func feature=pack type=resolver control=iteration dimension=2
//ff:what TestReaderPackInventoryReplay — raw and zstd sets of several packs and generations: replay equals State and the index, and each live Loc reads the path's bytes with the expected encoding
package pack

import (
	"bytes"
	"context"
	"reflect"
	"testing"
)

func TestReaderPackInventoryReplay(t *testing.T) {
	ctx := context.Background()
	for _, c := range []struct {
		codec Codec
		enc   byte
	}{{Stored{}, encRaw}, {&Zstd{DictSize: -1}, encZstd}} {
		out, docs := inventorySet(t, c.codec)
		r := opened(t, out)
		if n := len(r.DataPacks()); n < 3 {
			t.Fatal("want several data packs:", n)
		}
		st, err := r.State(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if got := replayed(t, r); !reflect.DeepEqual(got, st) {
			t.Fatalf("replay %v, State %v", got, st)
		}
		idx, err := r.entries(ctx)
		if err != nil || sameState(st, idx) != nil {
			t.Fatal(err)
		}
		if len(st) != 12 {
			t.Fatal(len(st), st)
		}
		for p, l := range st {
			b, err := r.Record(ctx, l)
			if err != nil || l.Enc != c.enc || !bytes.Equal(b, docs[p]) {
				t.Fatal(p, l, err)
			}
		}
	}
}
