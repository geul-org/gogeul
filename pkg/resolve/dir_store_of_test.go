//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirStoreOf — the given Store (over the loose files when a Root is given too), else a directory store at <Root>/.well-known/geul; publishing and reading work over a non-directory Store
package resolve

import (
	"context"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestDirStoreOf(t *testing.T) {
	if s, ok := (Dir{Root: "r"}).store().(dirStore); !ok || s.root != filepath.Join("r", ".well-known", "geul") {
		t.Fatal(s)
	}
	m := mapStore{fstest.MapFS{}}
	d := Dir{Store: m}
	if _, ok := d.store().(mapStore); !ok {
		t.Fatal("given store")
	}
	if o, ok := (Dir{Root: "r", Store: m}).store().(overlay); !ok || o.loose.root != filepath.Join("r", ".well-known", "geul") {
		t.Fatal("a Store with a Root")
	} else if _, ok := o.stored.(mapStore); !ok {
		t.Fatal("the stored side")
	}
	if err := d.write("0x0001/a.geul", []byte("x")); err != nil {
		t.Fatal(err)
	}
	if b, err := d.Fetch(context.Background(), "0x0001/a.geul"); err != nil || string(b) != "x" {
		t.Fatal(err)
	}
	if err := d.removeCopies([][32]byte{{1}}, nil); err != nil {
		t.Fatal(err)
	}
}
