//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDocs — <I>/docs/<V>.geul of a zero-padded Issuer; another Issuer's list is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocs(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := Docs(ctx, f.d, f.pad, 1); err != nil || !bytes.Equal(b, f.docs) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := Docs(ctx, f.d, []core.Word{0, 1}, Latest); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
