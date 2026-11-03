//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriteRefuses — a path outside the stored 공표 주소 forms (latest.geul, old h/ and documents.geul, names) is never stored
package pack

import (
	"context"
	"path/filepath"
	"testing"
)

func TestWriteRefuses(t *testing.T) {
	ctx := context.Background()
	for _, p := range []string{"0x00000000/docs/0x0000000000000008/latest.geul", "issuers/latest.geul", "issuers.geul",
		"0x0000/documents.geul", "h/0x" + "ab" + "00000000000000000000000000000000000000000000000000000000000000.geul", "0x00000000/keys/a.geul", "x.geul"} {
		out := filepath.Join(t.TempDir(), "p")
		if _, err := Write(ctx, out, Batch{Put: []string{p}, Get: memGet(map[string][]byte{p: {1}})}, BuildOptions{}); err == nil {
			t.Fatalf("%s stored", p)
		}
	}
}
