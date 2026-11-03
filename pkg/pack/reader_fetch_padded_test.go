//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderFetchPadded — the same document under an 8, 16 or 64-digit Issuer, a hashes/ copy only under its own Issuer, latest.geul not stored
package pack

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
)

func TestReaderFetchPadded(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	want := fixtureFile(t, "0x00000000/docs/0x0000000000000008/0x00000001.geul")
	for _, is := range []string{"0x00000000", "0x0000000000000000", "0x" + strings.Repeat("0", 64)} {
		got, err := r.Fetch(ctx, is+"/docs/0x0000000000000008/0x00000001.geul")
		if err != nil || !bytes.Equal(got, want) {
			t.Fatal(is, err)
		}
	}
	h := docHPath(want)
	if got, err := r.Fetch(ctx, h); err != nil || !bytes.Equal(got, want) {
		t.Fatal(h, err)
	}
	for _, p := range []string{strings.Replace(h, "0x00000000/", "0x00000002/", 1), "0x00000000/docs/0x0000000000000008/latest.geul", "issuers.geul"} {
		if _, err := r.Fetch(ctx, p); !errors.Is(err, ErrNotFound) {
			t.Fatal(p, err)
		}
	}
}
