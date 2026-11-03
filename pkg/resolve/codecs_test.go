//ff:func feature=resolve type=resolver control=sequence
//ff:what TestCodecs — <I>/codecs/latest.geul is the newest codec list; another Issuer's is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCodecs(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := Codecs(ctx, f.d, f.g, Latest); err != nil || !bytes.Equal(b, f.codecs) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := Codecs(ctx, f.d, []core.Word{0, 1}, Latest); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
