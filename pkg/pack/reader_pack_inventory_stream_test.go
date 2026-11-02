//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderPackInventoryStream — a Stream set (zstd with a dictionary, as Wikidata is written): replay equals State and each document's Loc reads its original bytes
package pack

import (
	"bytes"
	"context"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReaderPackInventoryStream(t *testing.T) {
	ctx := context.Background()
	out := filepath.Join(t.TempDir(), "set")
	s, err := NewStream(out, streamTestOptions())
	if err != nil {
		t.Fatal(err)
	}
	streamTestPut(t, s, 1, 40)
	res, err := s.Close()
	if err != nil || res.Dict == 0 {
		t.Fatalf("%+v %v", res, err)
	}
	r := opened(t, out)
	st, err := r.State(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if got := replayed(t, r); len(r.DataPacks()) < 2 || !reflect.DeepEqual(got, st) {
		t.Fatal(len(r.DataPacks()), len(got), len(st))
	}
	steps := inventorySteps(t, r)
	if len(steps) != 80 {
		t.Fatal(len(steps))
	}
	for i, e := range steps {
		b, err := r.Record(ctx, e.Loc)
		want := streamTestDoc(uint64(i/2+1), uint32(i%2))
		if err != nil || !e.Hashed || e.Deleted || e.Loc.Enc != encZstd || !bytes.Equal(b, want) {
			t.Fatalf("%d %+v %v", i, e, err)
		}
	}
}
