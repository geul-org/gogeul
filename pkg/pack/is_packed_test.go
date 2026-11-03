//ff:func feature=pack type=resolver control=sequence
//ff:what TestIsPacked — a top latest.txt or an Issuer folder with one; plain trees and other folders are not
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsPacked(t *testing.T) {
	if IsPacked(filepath.Join("testdata", "archive", ".well-known", "geul")) || IsPacked(t.TempDir()) {
		t.Fatal("plain")
	}
	top := t.TempDir()
	os.WriteFile(filepath.Join(top, latestFile), nil, 0o644)
	sub := t.TempDir()
	os.MkdirAll(filepath.Join(sub, "0x00000000"), 0o755)
	os.WriteFile(filepath.Join(sub, "0x00000000", latestFile), nil, 0o644)
	other := t.TempDir()
	os.MkdirAll(filepath.Join(other, "x"), 0o755)
	os.WriteFile(filepath.Join(other, "x", latestFile), nil, 0o644)
	if !IsPacked(top) || !IsPacked(sub) || IsPacked(other) {
		t.Fatal("packed")
	}
}
