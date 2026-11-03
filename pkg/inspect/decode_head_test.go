//ff:func feature=inspect type=decoder control=sequence
//ff:what TestDecodeHead — every Meta becomes an entry in order and the PROFILE text is returned
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDecodeHead(t *testing.T) {
	vw, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("p-1.0")))
	head := []core.Record{
		{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x0418, 0, 2, 0, 0, 0, 2, 0, 0, 3}},
		{Kind: core.KindMeta, Codec: core.MetaProfile, Payload: vw},
		core.EncodeSpecVersion(core.SpecVersion),
	}
	entries, profile := decodeHead(head)
	imp, ok := entries[0].Data.(ImportDoc)
	sv, _ := entries[2].Data.(Value)
	if len(entries) != 3 || !ok || imp.ID != 3 || profile != "p-1.0" || sv.Uint != 1 {
		t.Fatalf("got %+v %q", entries, profile)
	}
}
