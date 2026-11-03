//ff:func feature=server type=helper control=iteration dimension=1
//ff:what TestDocNameMaps — both directions; the same bytes come from the cache, changed bytes are parsed again, a full cache is emptied; a missing file gives empty maps
package server

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDocNameMaps(t *testing.T) {
	root := namesFixture(t)
	m := docNameMaps(treeOf(root), "0x00000000")
	if m.byName["publication"] != "0x00000980" || m.byID["0x00000980"] != "publication" || len(m.byID) != 2 {
		t.Fatal(m)
	}
	if again := docNameMaps(treeOf(root), "0x00000000"); fmt.Sprintf("%p", again.byName) != fmt.Sprintf("%p", m.byName) {
		t.Fatal("same bytes parsed again")
	}
	file := filepath.Join(root, ".well-known", "geul", "0x00000000", "docs", "0x00000001.geul")
	os.WriteFile(file, []byte{0}, 0o644)
	if m := docNameMaps(treeOf(root), "0x00000000"); len(m.byName) != 0 {
		t.Fatal("changed bytes not parsed again", m)
	}
	for i := 0; len(nameCache) < nameCacheMax; i++ {
		nameCache[string(rune(i))] = nameMaps{}
	}
	os.WriteFile(file, []byte{1}, 0o644)
	if docNameMaps(treeOf(root), "0x00000000"); len(nameCache) != 1 {
		t.Fatal("full cache kept", len(nameCache))
	}
	if m := docNameMaps(treeOf(root), "0x00000009"); m.byName == nil || m.byID == nil || len(m.byName) != 0 {
		t.Fatal("missing", m)
	}
}
