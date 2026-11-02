//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderRecord — raw and zstd records; a bad location fails
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderRecord(t *testing.T) {
	z := zstdFrame([]byte("zz"))
	obj, o := objBytes(kindData, append([]byte("ab"), z...), nil)
	r, ctx := newReader(memStore{objectName(o): obj}, Options{}, &manifest{objects: []object{o}}, ""), context.Background()
	if b, err := r.Record(ctx, Loc{Off: 64, Len: 2}); err != nil || string(b) != "ab" {
		t.Fatal(err)
	}
	if b, err := r.Record(ctx, Loc{Off: 66, Len: uint64(len(z)), Enc: encZstd}); err != nil || string(b) != "zz" {
		t.Fatal(err)
	}
	if _, err := r.Record(ctx, Loc{Obj: 3}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
