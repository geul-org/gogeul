//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDocVersions — <I>/docs/<D>/versions.geul; an unknown document is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestDocVersions(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := DocVersions(ctx, f.d, f.g, 0xffff000000000007); err != nil || !bytes.Equal(b, f.list) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := DocVersions(ctx, f.d, f.g, 0xffff000000000008); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
