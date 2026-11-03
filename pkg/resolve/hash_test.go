//ff:func feature=resolve type=resolver control=sequence
//ff:what TestHash — <I>/hashes/<H>.geul of the owning Issuer; another Issuer's hashes/ is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHash(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := Hash(ctx, f.d, f.g, f.hash); err != nil || !bytes.Equal(b, f.v2) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := Hash(ctx, f.d, []core.Word{0, 1}, f.hash); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
