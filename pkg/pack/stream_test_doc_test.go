//ff:func feature=pack type=writer control=sequence
//ff:what streamTestDoc — a small GEUL document of Issuer 0x00000001 for stream tests, distinct per document and slot, long enough to compress
package pack

import (
	"fmt"
	"slices"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// streamTestDoc builds document doc (version slot, 0 for the list) with a
// BIND text that repeats enough to compress.
func streamTestDoc(doc uint64, slot uint32) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocID([]core.Word{0, 1}, []core.Word{0x5100, 0, core.Word(doc >> 16), core.Word(doc)})
	text := strings.Repeat(fmt.Sprintf("entity %d slot %d label text; ", doc, slot), 12)
	bind, _ := core.EncodeBind(core.Bind{EdgeID: 1, Subject: core.EncodeLocalRef(1), Name: []byte(text)})
	return core.PutWords(core.WrapDocument(slices.Concat(enc(core.EncodeSpecVersion(1)), enc(id), []core.Word{0x1C30}, enc(bind), []core.Word{0x1C40})))
}
