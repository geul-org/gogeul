//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirDocuments — document then codec directories from the stored folders, without names; files and other entries skipped; a missing Issuer fails
package resolve

import (
	"os"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirDocuments(t *testing.T) {
	d := issuersFixture(t)
	os.MkdirAll(d.file("0x00000000/docs/x"), 0o755)
	os.MkdirAll(d.file("0x00000000/codecs/0x00000B80"), 0o755)
	os.WriteFile(d.file("0x00000000/docs/0x0000000000000009"), []byte{1}, 0o644)
	got, err := d.Documents([]core.Word{0, 0})
	if err != nil || len(got) != 2 || got[0].ID != "0xFFFF000000000007" || got[1].ID != "0x00000B80" || got[0].Names != nil {
		t.Fatalf("%+v %v", got, err)
	}
	if _, err := d.Documents([]core.Word{0, 1}); err == nil {
		t.Fatal("missing Issuer: want error")
	}
}
