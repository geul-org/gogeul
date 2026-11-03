//ff:func feature=cli type=helper control=sequence
//ff:what brokenDraft — test fixture: a draft that splits but breaks the spec (a FIELD Edge without its layout), so any save of it fails
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// brokenDraft writes `<dir>/broken.geul`: DOC-ID 0x00000000/…01, LABEL #1 and
// FIELD #2 holding only a Wikidata ID.
func brokenDraft(t *testing.T, dir string) string {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0, 0, 0, 1})
	q := core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{1}}
	label := core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: append(enc(core.EncodeLocalRef(1)), enc(q)...)}
	field := core.Record{Kind: core.KindEdge, Codec: core.EdgeField, Payload: append(enc(core.EncodeLocalRef(2)), enc(q)...)}
	b, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2), id}, []core.Record{label, field})
	p := filepath.Join(dir, "broken.geul")
	os.WriteFile(p, b, 0o644)
	return p
}
