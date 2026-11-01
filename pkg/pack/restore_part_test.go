//ff:func feature=pack type=writer control=sequence
//ff:what TestRestorePart — the part gets the sealed object's bytes after the header up to the size, a zero header; a missing source, an uncreatable part and an unreadable source fail
package pack

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRestorePart(t *testing.T) {
	dir := t.TempDir()
	src, dst := filepath.Join(dir, "sealed"), filepath.Join(dir, "part")
	b := bytes.Repeat([]byte("abcdefghij"), 20)
	os.WriteFile(src, b, 0o644)
	f, err := restorePart(dst, src, 150)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	got, _ := os.ReadFile(dst)
	if len(got) != 150 || !bytes.Equal(got[:headerSize], make([]byte, headerSize)) || !bytes.Equal(got[headerSize:], b[headerSize:150]) {
		t.Fatalf("restored %d bytes", len(got))
	}
	if _, err := restorePart(dst, filepath.Join(dir, "missing"), 150); err == nil {
		t.Fatal("missing source: want error")
	}
	if _, err := restorePart(filepath.Join(dir, "no", "part"), src, 150); err == nil {
		t.Fatal("uncreatable part: want error")
	}
	if _, err := restorePart(dst, dir, 150); err == nil {
		t.Fatal("unreadable source: want error")
	}
}
