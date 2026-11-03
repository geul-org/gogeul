//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestEdgeIDToken — a LOCAL first part is #n; none, a raw part or another record is false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestEdgeIDToken(t *testing.T) {
	if s, ok := edgeIDToken([]inspect.Part{{Words: []core.Word{0x0051, 4}, Record: true}}); !ok || s != "#4" {
		t.Fatal(s)
	}
	for _, ps := range [][]inspect.Part{nil, {{Words: []core.Word{4}}}, {{Words: []core.Word{0x0141, 4}, Record: true}}} {
		if _, ok := edgeIDToken(ps); ok {
			t.Fatalf("%+v: want false", ps)
		}
	}
}
