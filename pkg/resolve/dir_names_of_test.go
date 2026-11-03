//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirNamesOf — a missing index is empty; odd bytes and a broken document fail; a written index gives its names
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirNamesOf(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	if h, b, n, err := d.namesOf(namesPath(geulIssuer)); err != nil || h != nil || b != nil || len(n) != 0 {
		t.Fatalf("missing: %v", err)
	}
	for _, x := range [][]byte{{1}, {0, 5}} {
		d.write(namesPath(geulIssuer), x)
		if _, _, _, err := d.namesOf(namesPath(geulIssuer)); err == nil {
			t.Fatal("bad index: want error")
		}
	}
	subj, _ := core.EncodeDocPointer([]core.Word{0, 0}, []core.Word{0, 0, 0, 7}, []core.Word{0, 0})
	b, _ := namesDoc(nil, nil, core.Bind{EdgeID: 1, Subject: subj, Name: []byte("seven")})
	d.write(namesPath(geulIssuer), b)
	if _, _, n, err := d.namesOf(namesPath(geulIssuer)); err != nil || n["seven"].Doc != "0x0000000000000007" {
		t.Fatalf("%v %v", n, err)
	}
	if _, _, _, err := (Dir{Root: "/dev/null/x"}).namesOf("a"); err == nil {
		t.Fatal("unreadable root: want error")
	}
}
