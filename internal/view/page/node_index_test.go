//ff:func feature=page type=helper control=sequence
//ff:what TestNodeIndex — a new record gets a new row; the same words reuse it
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestNodeIndex(t *testing.T) {
	var rows []Row
	at := map[string]int{}
	a := inspect.Entry{Record: core.EncodeLocalRef(7), Words: []core.Word{0x0051, 7}}
	b := inspect.Entry{Record: core.EncodeLocalRef(8), Words: []core.Word{0x0051, 8}}
	if nodeIndex(&rows, at, a) != 0 || nodeIndex(&rows, at, b) != 1 || nodeIndex(&rows, at, a) != 0 || len(rows) != 2 {
		t.Fatalf("got %+v", rows)
	}
}
