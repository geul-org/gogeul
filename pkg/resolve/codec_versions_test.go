//ff:func feature=resolve type=resolver control=sequence
//ff:what TestCodecVersions — <I>/codecs/<C>/versions.geul; an unknown codec is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestCodecVersions(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := CodecVersions(ctx, f.d, f.pad, 0x0980); err != nil || !bytes.Equal(b, f.clist) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := CodecVersions(ctx, f.d, f.g, 0x0981); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
