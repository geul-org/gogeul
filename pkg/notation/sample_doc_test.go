//ff:func feature=notation type=helper control=iteration dimension=1
//ff:what sampleDoc — test fixture: a document with SPEC-VERSION, LABEL #1 Q42 "Douglas Adams"@en and TRIPLE #2 P31 Q42 Q5 rank 1
package notation

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// sampleDoc encodes a small Wikidata-like document without codec specs.
func sampleDoc() []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	q := func(c uint16, n core.Word) core.Record {
		return core.Record{Kind: core.KindReference, Codec: c, Payload: []core.Word{n}}
	}
	lang := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en"))
	label := core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: slices.Concat(enc(core.EncodeLocalRef(1)), enc(q(core.CodecWikidataQID, 42)), enc(core.EncodeLanguageText(lang, []byte("Douglas Adams"))))}
	triple := core.Record{Kind: core.KindEdge, Codec: 0x0A, Payload: append(slices.Concat(enc(core.EncodeLocalRef(2)), enc(q(core.CodecWikidataPID, 31)), enc(q(core.CodecWikidataQID, 42)), enc(q(core.CodecWikidataQID, 5))), 1)}
	in := enc(core.EncodeSpecVersion(2))
	in = append(in, 0x1C30)
	for _, r := range []core.Record{label, triple} {
		in = append(in, enc(r)...)
	}
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
