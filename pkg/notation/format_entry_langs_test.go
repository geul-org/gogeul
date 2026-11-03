//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatEntryLangs — a LABEL whose language is a LANGUAGE-SHORT-CODE shows the registry tag after @@ when the entry carries the registry, the hex code without it
package notation

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestFormatEntryLangs(t *testing.T) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	short := core.Record{Kind: core.KindValue, Codec: core.CodecLanguageShortCode, Payload: []core.Word{0x0784}}
	q42 := core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{42}}
	label := core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: slices.Concat(enc(core.EncodeLocalRef(1)), enc(q42), enc(core.EncodeLanguageText(short, []byte("Douglas Adams"))))}
	doc, err := inspect.Decode(core.PutWords(core.WrapDocument(slices.Concat(enc(core.EncodeSpecVersion(1)), []core.Word{0x1C30}, enc(label), []core.Word{0x1C40}))))
	if err != nil {
		t.Fatal(err)
	}
	e := doc.Body[0]
	if got := FormatEntry(e); got != `#1 LABEL Q42 "Douglas Adams"@@0x0784` {
		t.Fatalf("no registry: %s", got)
	}
	e.Langs = map[string]string{"0x0784": "en"}
	if got := FormatEntry(e); got != `#1 LABEL Q42 "Douglas Adams"@@en` {
		t.Fatalf("registry: %s", got)
	}
}
