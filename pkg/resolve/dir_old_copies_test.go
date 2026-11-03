//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirOldCopies — a two-list chain reaches both lists and both versions; no list, a missing parent and a broken parent stop the chain
package resolve

import (
	"os"
	"testing"
)

func TestDirOldCopies(t *testing.T) {
	d := sampleTree(t)
	k := sampleKey(1, 0).DocKey()
	lg, _ := d.readList(k, k.Doc)
	if hs := d.oldCopies(lg, k); len(hs) != 5 {
		t.Fatalf("chain: %d hashes", len(hs))
	}
	if hs := d.oldCopies(versionLog{}, k); hs != nil {
		t.Fatalf("no list: %v", hs)
	}
	os.WriteFile(d.file(hashPath(geulIssuer, lg.Parent)), []byte{1}, 0o644)
	if hs := d.oldCopies(lg, k); len(hs) != 3 {
		t.Fatalf("broken parent: %d", len(hs))
	}
	os.Remove(d.file(hashPath(geulIssuer, lg.Parent)))
	if hs := d.oldCopies(lg, k); len(hs) != 3 {
		t.Fatalf("missing parent: %d", len(hs))
	}
}
