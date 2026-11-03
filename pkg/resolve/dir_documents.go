//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.Documents — the documents an archive holds under one Issuer, with their names in binding order
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// Documents lists the document directories of issuer (`<Issuer>/docs/<문서 ID>`,
// canonical 4-word IDs in ID order) with the names its document list (`docs/`) binds to each, in BIND
// order so the last is the main name.
func (d Dir) Documents(issuer []core.Word) ([]Named, error) {
	is := core.FormatID(issuer)
	ents, err := d.store().ReadDir(is + "/" + docsDir)
	if err != nil {
		return nil, err
	}
	_, body, _, err := d.namesOf(namesPath(is))
	if err != nil {
		return nil, err
	}
	byDoc := bindOrder(body)
	var out []Named
	for _, e := range ents {
		if _, ok := parseHexID(e.Name(), 4); ok && e.IsDir() {
			out = append(out, Named{ID: e.Name(), Names: byDoc[e.Name()]})
		}
	}
	return out, nil
}
