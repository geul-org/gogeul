//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirPublish — sequential versions, PARENT chain, idempotent republish, list and hash files written; bad, unnamed, exhausted and broken-list inputs fail
package resolve

import (
	"context"
	"math"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestDirPublish(t *testing.T) {
	d := sampleTree(t)
	k := sampleKey(1, 0).DocKey()
	lg, err := d.readList(k, k.Doc)
	if err != nil || len(lg.Entries) != 2 || lg.Entries[0].Version != 1 || lg.Entries[1].Version != 2 {
		t.Fatalf("list: %+v %v", lg, err)
	}
	b, _ := d.Fetch(context.Background(), listPath(k))
	list, _ := inspect.Decode(b)
	if list.Head[1].Name != "PARENT" {
		t.Fatalf("second list has no PARENT: %+v", list.Head)
	}
	if e, err := d.Publish(sampleDoc(1, "A v1")); err != nil || e.Version != 1 {
		t.Fatalf("republish: %+v %v", e, err)
	}
	if b, err := d.Fetch(context.Background(), hashPath(geulIssuer, lg.Entries[1].Hash)); err != nil || len(b) == 0 {
		t.Fatalf("hash path: %v", err)
	}
	if b, err := d.Fetch(context.Background(), hashPath(geulIssuer, lg.Hash)); err != nil || len(b) == 0 {
		t.Fatalf("list kept under h/: %v", err)
	}
	if _, err := d.Publish([]byte{1}); err == nil {
		t.Fatal("odd bytes: want error")
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	listID, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0xffff, 0xffff, 0xffff, 0xfffe})
	if _, err := d.Publish(core.PutWords(core.WrapDocument(append(enc(listID), 0x1C30, 0x1C40)))); err == nil {
		t.Fatal("reserved list ID: want error")
	}
	if _, err := d.Publish(core.PutWords(core.WrapDocument([]core.Word{0x1C30}))); err == nil {
		t.Fatal("no END: want error")
	}
	if _, err := d.Publish(core.PutWords(core.WrapDocument([]core.Word{0x1C30, 0x1C40}))); err == nil {
		t.Fatal("unnamed: want error")
	}
	if err := d.write(listPath(sampleKey(4, 0).DocKey()), listDocOf(4, VersionEntry{Version: math.MaxUint32})); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Publish(sampleDoc(4, "x")); err == nil {
		t.Fatal("exhausted: want error")
	}
	if err := d.write(listPath(sampleKey(5, 0).DocKey()), []byte("junk")); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Publish(sampleDoc(5, "x")); err == nil {
		t.Fatal("bad list: want error")
	}
}
