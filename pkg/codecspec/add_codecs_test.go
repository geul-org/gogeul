//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestAddCodecs — a record field gathers every target, a raw field only its first; unknown fields and other records are skipped
package codecspec

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAddCodecs(t *testing.T) {
	fields := map[uint64]*Field{2: {}, 3: {Width: 4}}
	codec := func(id, f uint64, t core.Word) core.Record {
		return core.EncodeCodecDecl(core.CodecDecl{EdgeID: id, Field: f, Target: t})
	}
	addCodecs([]core.Record{codec(4, 2, 0x0580), codec(5, 2, 0x0050), codec(6, 3, 0x0430), codec(7, 3, 0x0410), codec(8, 9, 1), core.EncodeUintValue(1)}, fields)
	if !slices.Equal(fields[2].Codecs, []core.Word{0x0580, 0x0050}) || !slices.Equal(fields[3].Codecs, []core.Word{0x0430}) {
		t.Fatalf("%v %v", fields[2].Codecs, fields[3].Codecs)
	}
}
