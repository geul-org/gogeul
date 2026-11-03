//ff:func feature=notation type=decoder control=sequence
//ff:what TestDocAddress — canonical hex, @version only when not 0
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocAddress(t *testing.T) {
	if s := docAddress([]core.Word{0, 0}, []core.Word{0, 0, 0, 1}, 0); s != "0x00000000/0x0000000000000001" {
		t.Fatal(s)
	}
	if s := docAddress([]core.Word{0, 0}, []core.Word{0, 0, 0, 1}, 3); s != "0x00000000/0x0000000000000001@3" {
		t.Fatal(s)
	}
}
