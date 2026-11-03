//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirFile — paths map under .well-known/geul
package resolve

import (
	"path/filepath"
	"testing"
)

func TestDirFile(t *testing.T) {
	if f := (Dir{Root: "/r"}).file("0x00000000/docs/0x0000000000000001/versions.geul"); f != filepath.FromSlash("/r/.well-known/geul/0x00000000/docs/0x0000000000000001/versions.geul") {
		t.Fatalf("got %q", f)
	}
}
