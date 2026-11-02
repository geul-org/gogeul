//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderFetch — every fixture path reads back its file's bytes through P, H and named; others are NotFound
package pack

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestReaderFetch(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	paths, _ := listTree("testdata/archive")
	for _, p := range paths {
		b, err := r.Fetch(ctx, p)
		if err != nil || !bytes.Equal(b, fixtureFile(t, p)) {
			t.Fatal(p, err)
		}
	}
	for _, p := range []string{"0x0000/0x0000000000000009/versions.geul", "h/0x" + string(bytes.Repeat([]byte("0"), 64)) + ".geul", "none.geul"} {
		if _, err := r.Fetch(ctx, p); !errors.Is(err, ErrNotFound) {
			t.Fatal(p, err)
		}
	}
}
