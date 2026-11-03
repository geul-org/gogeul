//ff:func feature=resolve type=resolver control=sequence
//ff:what TestIssuers — issuers/latest.geul is the newest Issuer list; a version not published is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestIssuers(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := Issuers(ctx, f.d, Latest); err != nil || !bytes.Equal(b, f.issuers) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := Issuers(ctx, f.d, 2); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
