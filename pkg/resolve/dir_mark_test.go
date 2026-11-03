//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirMark — a mark appends one record with PARENT and keeps NEXT-LOCAL-ID; unpublished, twice deprecated, deleted, an unwritable archive and a broken list fail
package resolve

import (
	"os"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirMark(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	issuer, doc := []core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}
	d.Publish(edgeDoc(7, "a"))
	if _, err := d.mark(issuer, doc, 1, core.EdgeDeprecation, 0); err != nil {
		t.Fatal(err)
	}
	lg, _ := d.readList(sampleKeyOf(7), sampleKeyOf(7).Doc)
	if len(lg.Records) != 2 || !lg.HasParent || lg.Next != 2 || !lg.Entries[0].Deprecated {
		t.Fatalf("%+v", lg)
	}
	if _, err := d.mark(issuer, doc, 5, core.EdgeDeprecation, 0); err == nil {
		t.Fatal("not published: want error")
	}
	if _, err := d.mark(issuer, doc, 1, core.EdgeDeprecation, 0); err == nil {
		t.Fatal("deprecated twice: want error")
	}
	os.Chmod(d.file("0x00000000/hashes"), 0o500)
	defer os.Chmod(d.file("0x00000000/hashes"), 0o755)
	if _, err := d.mark(issuer, doc, 1, core.EdgeTombstone, 1); err == nil {
		t.Fatal("unwritable: want error")
	}
	os.Chmod(d.file("0x00000000/hashes"), 0o755)
	d.mark(issuer, doc, 1, core.EdgeTombstone, 1)
	if _, err := d.mark(issuer, doc, 1, core.EdgeTombstone, 1); err == nil {
		t.Fatal("deleted: want error")
	}
	d.write(listPath(sampleKeyOf(7)), []byte{1})
	if _, err := d.mark(issuer, doc, 1, core.EdgeDeprecation, 0); err == nil {
		t.Fatal("broken list: want error")
	}
}
