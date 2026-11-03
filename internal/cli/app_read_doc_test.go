//ff:func feature=cli type=helper control=sequence
//ff:what TestAppReadDoc — a file as is, an address from the data folder; a bad address and an address without a data folder fail
package cli

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestAppReadDoc(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	file, err := os.ReadFile(draft)
	if err != nil {
		t.Fatal(err)
	}
	if b, err := h.app(io.Discard).readDoc(draft); err != nil || !bytes.Equal(b, file) {
		t.Fatal(err)
	}
	if b, err := h.app(io.Discard).readDoc(addr); err != nil || !bytes.Equal(b, file) {
		t.Fatal(err)
	}
	if _, err := h.app(io.Discard).readDoc("nobody/x"); err == nil {
		t.Fatal("bad address: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.app(io.Discard).readDoc("0x00000000/0x0000000000000001"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
