//ff:func feature=cli type=helper control=sequence
//ff:what TestAppCommit — nothing for a directory archive; staged pack writes become a generation; a failing commit is returned
package cli

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestAppCommit(t *testing.T) {
	if err := (&app{}).commit(); err != nil {
		t.Fatal(err)
	}
	wk := filepath.Join(t.TempDir(), ".well-known", "geul")
	arch, _ := pack.OpenDir(context.Background(), wk, packOptions())
	arch.WriteFile("0x00000002/docs/0x0000000000000001/versions.geul", []byte("x"))
	a := &app{arch: arch}
	if err := a.commit(); err != nil || arch.Pending() || !pack.IsPacked(wk) {
		t.Fatal(err)
	}
	arch.WriteFile("0x00000000/hashes/0x0000000000000000000000000000000000000000000000000000000000000000.geul", []byte("not that hash"))
	if err := a.commit(); err == nil {
		t.Fatal("bad h/ copy: want error")
	}
}
