//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderDocuments — distinct document IDs of one Issuer in order; unknown Issuer none; walk errors fail
package pack

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestReaderDocuments(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	got, err := r.Documents(ctx, "0x00000000")
	want := []string{"0x0000000000000008", "0x0000000000000845", "0x0000000000000860", "0x000000000000180E"}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatal(got, err)
	}
	if got, err := r.Documents(ctx, "0x0009"); err != nil || got != nil {
		t.Fatal(got, err)
	}
	r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
	if _, err := r.Documents(ctx, "0x00000000"); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
