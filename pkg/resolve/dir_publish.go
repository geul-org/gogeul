//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Publish — assign the next version ID to a document and append it to its version list document
package resolve

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// Publish stores the document b in the local tree as the publisher does
// (GEUL-ORG 공표 주소, edge/PUBLICATION): it identifies the document by its
// DOC-ID, gives it the next version ID, writes `<Issuer>/<문서 ID>/<버전 ID>.geul`
// and `h/<내용 해시>.geul`, and writes a new version list document
// (`<Issuer>/<문서 ID>/versions.geul`) that keeps the previous records, appends
// a PUBLICATION and points PARENT at the previous list, which stays under `h/`.
// Publishing content already in the list returns that entry and writes nothing.
func (d Dir) Publish(b []byte) (VersionEntry, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return VersionEntry{}, err
	}
	h, err := core.ContentHash(ws)
	if err != nil {
		return VersionEntry{}, err
	}
	doc, _ := inspect.Decode(b)
	id, ok := docIDOf(doc)
	if !ok {
		return VersionEntry{}, errors.New("geul/resolve: publishing needs a DOC-ID")
	}
	if len(id.Doc) == 4 && core.IsListDocID(wordsUint(id.Doc)) {
		return VersionEntry{}, errors.New("geul/resolve: a reserved list document ID is published only as a list (PublishList)")
	}
	return d.publishAt(Key{Issuer: id.Issuer, Doc: id.Doc}.DocKey(), id.Issuer, id.Doc, b, h)
}
