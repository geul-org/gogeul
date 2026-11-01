//ff:func feature=pack type=writer control=sequence
//ff:what TestCutPart — a part is hashed after the header and cut to the size; a closed file fails and is not returned
package pack

import (
	"bytes"
	"crypto/sha256"
	"os"
	"path/filepath"
	"testing"
)

func TestCutPart(t *testing.T) {
	path := filepath.Join(t.TempDir(), "part")
	b := bytes.Repeat([]byte("0123456789"), 10)
	os.WriteFile(path, b, 0o644)
	f, _ := os.OpenFile(path, os.O_RDWR, 0)
	g, h, err := cutPart(f, 80)
	want := sha256.Sum256(b[headerSize:80])
	if fi, _ := os.Stat(path); err != nil || g != f || fi.Size() != 80 || !bytes.Equal(h.Sum(nil), want[:]) {
		t.Fatalf("%v", err)
	}
	f.Close()
	if g, h, err := cutPart(f, 80); err == nil || g != nil || h != nil {
		t.Fatal("closed file: want error")
	}
}
