//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirOverwrite — a document with two versions becomes one version 1 with a one-entry list, no other version files and no old h/ copies; rerun is stable; an undeletable version file, bad input and an unwritable root fail
package resolve

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestDirOverwrite(t *testing.T) {
	d := sampleTree(t)
	k := sampleKey(1, 0).DocKey()
	b := sampleDoc(1, "A dev")
	before, _ := d.readList(k, k.Doc)
	olds := d.oldCopies(before, k)
	e, err := d.Overwrite(b)
	if err != nil || e.Version != 1 {
		t.Fatalf("%+v %v", e, err)
	}
	lg, err := d.readList(k, k.Doc)
	if err != nil || len(lg.Entries) != 1 || lg.Entries[0].Version != 1 || lg.Entries[0].Hash != e.Hash {
		t.Fatalf("list: %+v %v", lg, err)
	}
	if len(olds) < 4 {
		t.Fatalf("old copies %d", len(olds))
	}
	for _, h := range olds {
		if _, err := os.Stat(d.file(hashPath(geulIssuer, h))); !os.IsNotExist(err) {
			t.Fatalf("old copy %x left: %v", h[:4], err)
		}
	}
	if _, err := os.Stat(d.file(hashPath(geulIssuer, e.Hash))); err != nil {
		t.Fatalf("new copy: %v", err)
	}
	got, _ := d.Fetch(context.Background(), versionPath(k, 1))
	if string(got) != string(b) {
		t.Fatal("version 1 not replaced")
	}
	if _, err := os.Stat(d.file(versionPath(k, 2))); !os.IsNotExist(err) {
		t.Fatalf("version 2 left: %v", err)
	}
	lb, _ := d.Fetch(context.Background(), listPath(k))
	if list, _ := inspect.Decode(lb); len(list.Head) != 3 || list.Head[1].Name != "RECORDED-AT" || list.Head[2].Name != "NEXT-LOCAL-ID" {
		t.Fatalf("list HEAD (no PARENT, NEXT-LOCAL-ID last): %+v", list.Head)
	}
	again, _ := d.Fetch(context.Background(), listPath(k))
	if _, err := d.Overwrite(b); err != nil {
		t.Fatal(err)
	}
	if after, _ := d.Fetch(context.Background(), listPath(k)); string(after) != string(again) {
		t.Fatal("rerun changed the list")
	}
	for _, bad := range [][]byte{{1}, core.PutWords(core.WrapDocument([]core.Word{0x1C30})), core.PutWords(core.WrapDocument([]core.Word{0x1C30, 0x1C40}))} {
		if _, err := d.Overwrite(bad); err == nil {
			t.Fatal("bad document: want error")
		}
	}
	stuck := filepath.Join(d.file(versionPath(k, 3)), "x")
	os.MkdirAll(stuck, 0o755)
	if _, err := d.Overwrite(sampleDoc(1, "A dev 2")); err == nil {
		t.Fatal("undeletable version file: want error")
	}
	blocked := Dir{Root: filepath.Join(t.TempDir(), "file")}
	os.WriteFile(blocked.Root, nil, 0o644)
	if _, err := blocked.Overwrite(b); err == nil {
		t.Fatal("unwritable root: want error")
	}
}
