//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirCurrent — nothing before publishing; the newest version with NEXT-LOCAL-ID; a stored diff is rebuilt; a list with no live version errors
package resolve

import (
	"bytes"
	"os"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/delta"
)

func TestDirCurrent(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	issuer, doc := []core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}
	if b, n, err := d.Current(issuer, doc); b != nil || n != 0 || err != nil {
		t.Fatalf("unpublished: %v %d %v", b, n, err)
	}
	v1, v2 := edgeDoc(7, "a", "b"), edgeDoc(7, "a", "B", "c")
	d.Publish(v1)
	d.Publish(v2)
	if b, n, err := d.Current(issuer, doc); err != nil || !bytes.Equal(b, v2) || n != 4 {
		t.Fatalf("v2: %d %v", n, err)
	}
	diff, err := delta.Diff(v1, v2, 1)
	if err != nil {
		t.Fatal(err)
	}
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	os.WriteFile(d.file(versionPath(k, 2)), diff, 0o644)
	if b, _, err := d.Current(issuer, doc); err != nil || !bytes.Equal(b, v2) {
		t.Fatalf("from diff: %v", err)
	}
	gone := Dir{Root: t.TempDir()}
	gone.Publish(v1)
	lg, _ := gone.readList(k, k.Doc)
	tomb := core.EncodePublication(core.Publication{Codec: core.EdgeTombstone, EdgeID: 2, Issuer: issuer, Doc: doc, Version: 1})
	gone.write(listPath(k), encodeVersionList(append(lg.Records, tomb), nil, 0, 0))
	if _, _, err := gone.Current(issuer, doc); err == nil {
		t.Fatal("no live version: want error")
	}
	bad := Dir{Root: t.TempDir()}
	bad.write(listPath(k), []byte{1})
	if _, _, err := bad.Current(issuer, doc); err == nil {
		t.Fatal("broken list: want error")
	}
}
