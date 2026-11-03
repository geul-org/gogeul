//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestFirstVersionList — the list equals what Overwrite writes for a new document but for RECORDED-AT; odd bytes, a non-document and a document without DOC-ID fail
package resolve

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFirstVersionList(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	b := edgeDoc(7, "a", "b")
	if _, err := d.Overwrite(b); err != nil {
		t.Fatal(err)
	}
	k := sampleKeyOf(7)
	want, err := os.ReadFile(filepath.Join(d.Root, ".well-known/geul", listPath(k)))
	if err != nil {
		t.Fatal(err)
	}
	h, _ := d.Log([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7})
	got, err := FirstVersionList(b, h.Entries[0].At.Unix())
	if err != nil || !bytes.Equal(got, want) {
		t.Fatalf("list differs: %v", err)
	}
	noID := core.PutWords(core.WrapDocument([]core.Word{0x1C30, 0x1C40}))
	for _, bad := range [][]byte{{1}, {0x12, 0x34}, noID} {
		if _, err := FirstVersionList(bad, 0); err == nil {
			t.Fatalf("%x: want error", bad)
		}
	}
}
