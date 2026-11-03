//ff:func feature=resolve type=resolver control=sequence
//ff:what TestIssuersVersions — issuers/versions.geul; none in an empty tree
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestIssuersVersions(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := IssuersVersions(ctx, f.d); err != nil || !bytes.Equal(b, f.ilist) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := IssuersVersions(ctx, Dir{Root: t.TempDir()}); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
