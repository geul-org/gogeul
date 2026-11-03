//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterDeleteEntry — P deletion entry, deleted hash, named deletion with a path table entry
package pack

import (
	"strings"
	"testing"
)

func TestWriterDeleteEntry(t *testing.T) {
	w := &writer{inum: map[string]uint32{"0x00000000": 0}}
	w.deleteEntry("0x00000000/docs/0x0000000000000008/versions.geul")
	w.deleteEntry("0x00000000/hashes/0x" + strings.Repeat("00", 32) + ".geul")
	w.deleteEntry("a.geul")
	w.deleteEntry("0x00000000/codecs/0x00000abc/0x0000000a.geul")
	if len(w.inv.entries) != 3 || w.inv.entries[0].kind != invPDel || w.inv.entries[0].doc != 8 || w.inv.entries[1].kind != invNamedDel || len(w.inv.deleted) != 1 || w.inv.paths[0] != "a.geul" {
		t.Fatal(w.inv)
	}
	if w.inv.entries[2].kind != invNamedDel || w.inv.paths[1] != "0x00000000/codecs/0x00000ABC/0x0000000A.geul" {
		t.Fatal("a named path is deleted under its stored spelling:", w.inv.paths)
	}
}
