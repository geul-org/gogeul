//ff:func feature=resolve type=resolver control=sequence
//ff:what TestCodecsVersions — <I>/codecs/versions.geul; another Issuer's is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCodecsVersions(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := CodecsVersions(ctx, f.d, f.pad); err != nil || !bytes.Equal(b, f.cslist) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := CodecsVersions(ctx, f.d, []core.Word{0, 1}); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
