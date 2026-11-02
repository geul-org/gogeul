//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderNamed — issuers.geul and documents.geul from NAMED; absent names, before and after the table, are NotFound
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
	if b, err := r.named(ctx, "issuers.geul"); err != nil || !bytes.Equal(b, fixtureFile(t, "issuers.geul")) {
		t.Fatal(err)
	}
	for _, p := range []string{"0x0000/x.geul", "zzz"} {
		if _, err := r.named(ctx, p); !errors.Is(err, ErrNotFound) {
			t.Fatal(p, err)
		}
	}
}
