//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestIssuerFolders — canonical Issuer folders only; none for an absent folder; a file for a folder fails
package resolve

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestIssuerFolders(t *testing.T) {
	dir := t.TempDir()
	for _, n := range []string{"0x00000002", "0x00000001", "h", "0xZZZZ"} {
		os.Mkdir(filepath.Join(dir, n), 0o755)
	}
	os.WriteFile(filepath.Join(dir, "0x00000003"), nil, 0o644)
	if got, err := issuerFolders(dir); err != nil || !slices.Equal(got, []string{"0x00000001", "0x00000002"}) {
		t.Fatal(got, err)
	}
	if got, err := issuerFolders(filepath.Join(dir, "absent")); err != nil || got != nil {
		t.Fatal(got, err)
	}
	if _, err := issuerFolders(filepath.Join(dir, "0x00000003")); err == nil {
		t.Fatal("a file")
	}
}
