//ff:func feature=cli type=helper control=sequence
//ff:what TestAppFindDoc — a file as is, a 공표 주소 from the data folder; bad local keys and no data folder fail
package cli

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestAppFindDoc(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	file, err := os.ReadFile(draft)
	if err != nil {
		t.Fatal(err)
	}
	if b, err := h.app(io.Discard).findDoc(draft); err != nil || !bytes.Equal(b, file) {
		t.Fatal(err)
	}
	if b, err := h.app(io.Discard).findDoc(addr + ".geul"); err != nil || !bytes.Equal(b, file) {
		t.Fatal(err)
	}
	cwd := h.cwd
	h.cwd = t.TempDir()
	if _, err := h.app(io.Discard).findDoc(addr); err == nil {
		t.Fatal("no data folder: want error")
	}
	h.cwd = cwd
	keys := filepath.Join(h.conf, "geul", "keys")
	os.RemoveAll(keys)
	os.MkdirAll(filepath.Dir(keys), 0o755)
	os.WriteFile(keys, nil, 0o644)
	if _, err := h.app(io.Discard).findDoc(addr); err == nil {
		t.Fatal("unreadable keys folder: want error")
	}
}
