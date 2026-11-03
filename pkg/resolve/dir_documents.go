//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Documents — the documents an archive holds under one Issuer: document IDs, then codec IDs, from the stored folders
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// Documents lists the document directories of issuer: `<Issuer>/docs/<문서 ID>`
// (canonical 4-word IDs in ID order), then its codecs `<Issuer>/codecs/<코덱 ID>`
// (2-word IDs). They are read from the stored folders (a pack set's
// inventory), not from a list document: there is none and documents have no
// names (§1.1 101), so Names is empty. An Issuer without documents (codecs
// only) lists none; a missing Issuer is an error.
func (d Dir) Documents(issuer []core.Word) ([]Named, error) {
	is := core.FormatID(issuer)
	if _, err := d.store().ReadDir(is); err != nil {
		return nil, err
	}
	out := d.namedIn(is+"/"+docsDir, 4)
	out = append(out, d.namedIn(is+"/"+codecsDir, 2)...)
	return out, nil
}
