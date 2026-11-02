//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestAppLocalIssuer — an Issuer part naming a local key becomes its ID, anything else stays; an unreadable keys folder fails
package cli

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppLocalIssuer(t *testing.T) {
	h := newHarness(t)
	h.must("issuer", "new", "--name", "me")
	a := h.app(io.Discard)
	got, err := a.localIssuer("me/doc")
	if err != nil || !strings.HasPrefix(got, "0x") || !strings.HasSuffix(got, "/doc") {
		t.Fatal(got, err)
	}
	for _, s := range []string{"other/doc", "0x0000/doc"} {
		if got, err := a.localIssuer(s); err != nil || got != s {
			t.Fatal(s, got, err)
		}
	}
	keys := filepath.Join(h.conf, "geul", "keys")
	os.RemoveAll(keys)
	os.WriteFile(keys, nil, 0o644)
	if _, err := a.localIssuer("me/doc"); err == nil {
		t.Fatal("unreadable keys folder: want error")
	}
}
