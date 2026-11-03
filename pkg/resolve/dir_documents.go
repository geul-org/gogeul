//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Documents — the documents an archive holds under one Issuer, with their names in binding order
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// Documents lists the document directories of issuer: `<Issuer>/docs/<문서 ID>`
// (canonical 4-word IDs in ID order), then its codecs `<Issuer>/codecs/<코덱 ID>`
// (2-word IDs), with the names its document list (`docs/`) binds to each, in BIND
// order so the last is the main name. An Issuer without documents (codecs
// only) lists none; a missing Issuer is an error.
func (d Dir) Documents(issuer []core.Word) ([]Named, error) {
	is := core.FormatID(issuer)
	if _, err := d.store().ReadDir(is); err != nil {
		return nil, err
	}
	_, body, _, err := d.namesOf(namesKey(is))
	if err != nil {
		return nil, err
	}
	byDoc := bindOrder(body)
	out := d.namedIn(is+"/"+docsDir, 4, byDoc)
	out = append(out, d.namedIn(is+"/"+codecsDir, 2, byDoc)...)
	return out, nil
}
