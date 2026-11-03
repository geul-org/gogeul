//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDoc — <I>/docs/<D>/latest.geul is version 2 for a padded Issuer; a 1-word Issuer is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDoc(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := Doc(ctx, f.d, f.pad, 0xffff000000000007, Latest); err != nil || !bytes.Equal(b, f.v2) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := Doc(ctx, f.d, []core.Word{0}, 0xffff000000000007, Latest); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
