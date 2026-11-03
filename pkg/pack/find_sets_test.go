//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestFindSets — the top set first, then Issuer sets in Dir order; a missing directory has none; an unreadable one fails
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindSets(t *testing.T) {
	wk := t.TempDir()
	for _, d := range []string{"", "0x00000001", "0x00000000", "zz"} {
		os.MkdirAll(filepath.Join(wk, d), 0o755)
		os.WriteFile(filepath.Join(wk, d, latestFile), nil, 0o644)
	}
	sets, err := findSets(wk)
	if err != nil || len(sets) != 3 || sets[0].name != "" || sets[1].name != "0x00000000" || sets[2].name != "0x00000001" {
		t.Fatal(sets, err)
	}
	if sets, err := findSets(filepath.Join(wk, "none")); err != nil || sets != nil {
		t.Fatal(sets, err)
	}
	f := filepath.Join(wk, "file")
	os.WriteFile(f, nil, 0o644)
	if _, err := findSets(f); err == nil {
		t.Fatal("not a directory")
	}
}
