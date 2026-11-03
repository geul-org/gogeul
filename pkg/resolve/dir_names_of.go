//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.namesOf — an Issuer's document list split into HEAD and BODY with its name map; all empty when the index does not exist
package resolve

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// namesOf reads the newest version of the name list k (latest). A missing
// list gives nil records and an empty map; one that does not split is an error.
func (d Dir) namesOf(k DocKey) ([]core.Record, []core.Record, map[string]DocKey, error) {
	b, err := d.latest(k)
	if errors.Is(err, ErrNotFound) {
		return nil, nil, map[string]DocKey{}, nil
	}
	if err != nil {
		return nil, nil, nil, err
	}
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, nil, nil, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return nil, nil, nil, err
	}
	names, _ := ParseDocumentNames(b)
	return doc.Head, doc.Body, names, nil
}
