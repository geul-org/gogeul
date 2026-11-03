//ff:func feature=delta type=helper control=sequence
//ff:what tripleOf — test fixture: a TRIPLE Edge #id with PID predicate, a subject record, a UTF8-TEXT object and rank normal
package delta

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// tripleOf is TRIPLE(#id, P pid, subject, "text", normal) for tests.
func tripleOf(id uint64, pid core.Word, subject core.Record, text string) core.Record {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	p := slices.Concat(enc(core.EncodeLocalRef(id)), []core.Word{0x0191, pid}, enc(subject), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(text))), []core.Word{0})
	return core.Record{Kind: core.KindEdge, Codec: edgeTriple, Payload: p}
}
