//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what checkAddrs — one storage form gives the wanted bytes at each address and ErrNotFound at each missing one
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

// checkAddrs reads every address of got and missing from f through FetchAddr.
func checkAddrs(t *testing.T, f form, got map[string][]byte, missing []string) {
	t.Helper()
	ctx := context.Background()
	for a, want := range got {
		addr, err := ParseAddr(a)
		if err != nil {
			t.Fatal(a, err)
		}
		if b, err := FetchAddr(ctx, f.dir, addr); err != nil || !bytes.Equal(b, want) {
			t.Fatalf("%s: %s: %d bytes %v", f.name, a, len(b), err)
		}
	}
	for _, a := range missing {
		addr, _ := ParseAddr(a)
		if _, err := FetchAddr(ctx, f.dir, addr); !errors.Is(err, ErrNotFound) {
			t.Fatalf("%s: %s: %v", f.name, a, err)
		}
	}
}
