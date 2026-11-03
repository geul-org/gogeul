//ff:func feature=cli type=helper control=sequence
//ff:what TestGetAddr — a file:// 공표 주소 reads its stored bytes (nil context allowed); an address outside the forms and one without a base fail
package cli

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestGetAddr(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, ".well-known", "geul", "issuers")
	os.MkdirAll(dir, 0o755)
	os.WriteFile(filepath.Join(dir, "versions.geul"), []byte{1, 2}, 0o644)
	if b, err := getAddr(nil, "file://"+root+"/.well-known/geul/issuers/versions.geul"); err != nil || len(b) != 2 {
		t.Fatal(b, err)
	}
	if _, err := getAddr(context.Background(), "file://"+root+"/.well-known/geul/issuers.geul"); err == nil {
		t.Fatal("old path: want error")
	}
	if _, err := getAddr(context.Background(), "issuers/versions.geul"); err == nil {
		t.Fatal("no base: want error")
	}
}
