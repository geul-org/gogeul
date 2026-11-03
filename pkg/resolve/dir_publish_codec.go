//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.PublishCodec — publish a direct codec's spec document after checking its DOC-ID names that codec
package resolve

import (
	"errors"
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// PublishCodec publishes the codec spec document b (edge/DECLARE 규격 문서)
// like any document (Publish). A spec document is the GEUL Issuer
// 0x00000000's codec codecDocID of the codec its one DECLARE targets (§3.13
// `0x00000000/codecs/<코덱 ID>/`), so b must hold exactly one DECLARE, of a
// direct codec, and a DOC-ID CODEC-POINTER naming that codec.
func (d Dir) PublishCodec(b []byte) (VersionEntry, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return VersionEntry{}, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return VersionEntry{}, err
	}
	target, ok := declareTarget(doc.Body)
	if !ok {
		return VersionEntry{}, errors.New("geul/resolve: a codec spec document needs exactly one DECLARE of a direct codec")
	}
	parsed, _ := inspect.Decode(b)
	if id, ok := docIDOf(parsed); !ok || !slices.Equal(id.Issuer, []core.Word{0, 0}) || !slices.Equal(id.Doc, codecDocID(target)) {
		return VersionEntry{}, errors.New("geul/resolve: a codec document needs DOC-ID CODEC-POINTER(0x00000000, 0x0000<codec TAG>, 0)")
	}
	return d.Publish(b)
}
