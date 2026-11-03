//ff:func feature=cli type=helper control=sequence
//ff:what TestAppKey — a local key name becomes its Issuer ID; IDs work without a data folder; a broken key fails
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppKey(t *testing.T) {
	h := newHarness(t)
	out := h.must("issuer", "new", "--name", "me")
	k, err := h.key("me/0x0000000000000001@2")
	if err != nil || !strings.HasSuffix(strings.TrimSpace(out), core.FormatID(k.Issuer)) || k.Version != 2 {
		t.Fatalf("%+v %v", k, err)
	}
	h.cwd = t.TempDir()
	if _, err := h.key("0x00000000/0x0000000000000001"); err != nil {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(h.conf, "geul", "keys", "x.pem"), []byte("x"), 0o600)
	if _, err := h.key("0x00000000/0x0000000000000001"); err == nil {
		t.Fatal("broken key: want error")
	}
}
