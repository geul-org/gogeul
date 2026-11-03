//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderNamed — the Issuer list from NAMED; absent names, before and after the table, are NotFound
package pack

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestReaderNamed(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	if b, err := r.named(ctx, "issuers/0x00000001.geul"); err != nil || !bytes.Equal(b, fixtureFile(t, "issuers/0x00000001.geul")) {
		t.Fatal(err)
	}
	for _, p := range []string{"0x00000000/x.geul", "zzz"} {
		if _, err := r.named(ctx, p); !errors.Is(err, ErrNotFound) {
			t.Fatal(p, err)
		}
	}
}
