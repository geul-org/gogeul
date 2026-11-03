//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirRemoveOtherVersions — version files other than 1 go, version 1 and the list stay; a missing folder is fine; an unlistable folder or undeletable file fails
package resolve

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDirRemoveOtherVersions(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	k := DocKey{Issuer: "0x00000002", Doc: "0x0000000000000001"}
	dir := d.file(docFolder(k))
	os.MkdirAll(dir, 0o755)
	for _, f := range []string{"0x00000001.geul", "0x00000002.geul", "versions.geul"} {
		os.WriteFile(filepath.Join(dir, f), []byte{0}, 0o644)
	}
	if err := d.removeOtherVersions(k); err != nil {
		t.Fatal(err)
	}
	left, _ := os.ReadDir(dir)
	if len(left) != 2 {
		t.Fatalf("left %d files", len(left))
	}
	if err := d.removeOtherVersions(DocKey{Issuer: "0x00000003", Doc: "0x0000000000000001"}); err != nil {
		t.Fatalf("missing folder: %v", err)
	}
	os.Mkdir(filepath.Join(dir, "0x00000003.geul"), 0o755)
	os.WriteFile(filepath.Join(dir, "0x00000003.geul", "x"), nil, 0o644)
	if err := d.removeOtherVersions(k); err == nil {
		t.Fatal("non-empty folder named like a version: want error")
	}
	f := DocKey{Issuer: "0x00000004", Doc: "0x0000000000000001"}
	os.MkdirAll(d.file(f.Issuer+"/docs"), 0o755)
	os.WriteFile(d.file(docFolder(f)), nil, 0o644)
	if err := d.removeOtherVersions(f); err == nil {
		t.Fatal("document folder is a file: want error")
	}
}
