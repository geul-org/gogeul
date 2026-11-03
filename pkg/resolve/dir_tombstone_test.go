//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirTombstone — a deleted version's file and h/ copy go; a version another is a diff on is kept; a second delete fails; an unremovable file errors
package resolve

import (
	"os"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/delta"
)

func TestDirTombstone(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	issuer, doc := []core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}
	k := sampleKeyOf(7)
	v1, v2 := edgeDoc(7, "a"), edgeDoc(7, "a", "b")
	e1, _ := d.Publish(v1)
	d.Publish(v2)
	diff, _ := delta.Diff(v1, v2, 1)
	os.WriteFile(d.file(versionPath(k, 2)), diff, 0o644)
	if err := d.Tombstone(issuer, doc, 1, 1); err == nil {
		t.Fatal("base of a diff: want error")
	}
	os.WriteFile(d.file(versionPath(k, 2)), v2, 0o644)
	if err := d.Tombstone(issuer, doc, 1, 1); err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{versionPath(k, 1), hashPath(geulIssuer, e1.Hash)} {
		if _, err := os.Stat(d.file(p)); !os.IsNotExist(err) {
			t.Fatalf("%s kept", p)
		}
	}
	if err := d.Tombstone(issuer, doc, 1, 1); err == nil {
		t.Fatal("twice: want error")
	}
	os.Remove(d.file(versionPath(k, 2)))
	os.MkdirAll(d.file(versionPath(k, 2))+"/x", 0o755)
	if err := d.Tombstone(issuer, doc, 2, 1); err == nil {
		t.Fatal("unremovable: want error")
	}
}
