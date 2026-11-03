//ff:func feature=resolve type=resolver control=sequence
//ff:what TestFetch — geul get's path: a file:// address reads the bytes; a missing address is ErrNotFound; an address without a base fails
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestFetch(t *testing.T) {
	ctx := context.Background()
	tree := formsTree(t)
	want, _ := tree.Find("0x00000000/docs/0xffff000000000007/latest.geul")
	a, err := ParseAddr("file://" + tree.Root + "/.well-known/geul/0x00000000/docs/0xffff000000000007/latest.geul")
	if err != nil {
		t.Fatal(err)
	}
	if b, err := Fetch(ctx, a); err != nil || !bytes.Equal(b, want) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	a, _ = ParseAddr("file://" + tree.Root + "/.well-known/geul/0x00000000/docs/0xffff000000000009/latest.geul")
	if _, err := Fetch(ctx, a); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	a, _ = ParseAddr("0x00000000/docs/0xffff000000000007/latest.geul")
	if _, err := Fetch(ctx, a); err == nil {
		t.Fatal("no base: want error")
	}
}
