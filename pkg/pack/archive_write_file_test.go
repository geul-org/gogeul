//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveWriteFile — staged bytes are copied, undo a staged removal and join the listing; invalid names and remote archives fail
package pack

import (
	"context"
	"errors"
	"io/fs"
	"testing"
)

func TestArchiveWriteFile(t *testing.T) {
	a, _ := testArchive(t)
	a.index()
	b := []byte("x")
	a.Remove("issuers/0x00000001.geul")
	if err := a.WriteFile("issuers/0x00000001.geul", b); err != nil || a.dels["issuers/0x00000001.geul"] {
		t.Fatal(err)
	}
	b[0] = 'y'
	if got, _ := a.ReadFile("issuers/0x00000001.geul"); string(got) != "x" {
		t.Fatal(string(got))
	}
	if err := a.WriteFile("0x0001/a.geul", b); err != nil || !hasPrefix(a.idx, "0x0001/") {
		t.Fatal(err)
	}
	if err := a.WriteFile(".", b); !errors.Is(err, fs.ErrInvalid) {
		t.Fatal(err)
	}
	_, out := built(t, small())
	r, _ := OpenRemote(context.Background(), FileStore{Root: out}, Options{})
	if err := r.WriteFile("a", b); err == nil {
		t.Fatal("remote")
	}
}
