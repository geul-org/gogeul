//ff:func feature=schema type=decoder control=sequence
//ff:what TestPartKind — the Kind of a record and whether it is a LOCAL Reference
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestPartKind(t *testing.T) {
	if k, local := partKind(enc(core.EncodeLocalRef(3))); k != core.KindReference || !local {
		t.Fatal(k, local)
	}
	if k, local := partKind([]core.Word{0x0141, 42}); k != core.KindReference || local {
		t.Fatal(k, local)
	}
	if k, local := partKind([]core.Word{0x0621, 7}); k != core.KindValue || local {
		t.Fatal(k, local)
	}
}
