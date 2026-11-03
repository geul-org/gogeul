//ff:func feature=inspect type=decoder control=sequence
//ff:what TestPartsLocal — the largest LOCAL among record parts; raw parts, broken records and other References are skipped
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestPartsLocal(t *testing.T) {
	local := func(id uint64) Part {
		ws, _ := core.EncodeRecord(core.EncodeLocalRef(id))
		return Part{Words: ws, Record: true}
	}
	ps := []Part{local(4), {Words: []core.Word{0x0051, 0x0063}}, {Words: []core.Word{0x0051}, Record: true}, {Words: []core.Word{0x0141, 0x0063}, Record: true}, local(2)}
	if got := partsLocal(ps); got != 4 {
		t.Fatalf("got %d", got)
	}
}
