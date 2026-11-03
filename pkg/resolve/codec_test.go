//ff:func feature=resolve type=resolver control=sequence
//ff:what TestCodec — <I>/codecs/<C>/<V>.geul; an unknown codec is ErrNotFound
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestCodec(t *testing.T) {
	ctx := context.Background()
	f := newTypedFixture(t)
	if b, err := Codec(ctx, f.d, f.g, 0x0980, 1); err != nil || !bytes.Equal(b, f.codec) {
		t.Fatalf("%d bytes %v", len(b), err)
	}
	if _, err := Codec(ctx, f.d, f.g, 0x0981, Latest); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
}
