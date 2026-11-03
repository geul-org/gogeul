//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeFuncRef — Alias REF and VerbCode decode; a missing field is an error
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDecodeFuncRef(t *testing.T) {
	pos := 0
	if f, err := decodeFuncRef([]core.Word{0x0051, 2, 0x7400}, &pos, testWire); err != nil || f.Alias != 2 || f.Verb != 0x7400 {
		t.Fatalf("got %+v %v", f, err)
	}
	for _, ws := range [][]core.Word{{}, {0x0051, 2}} {
		p := 0
		if _, err := decodeFuncRef(ws, &p, testWire); err == nil {
			t.Fatalf("%v: want error", ws)
		}
	}
}
