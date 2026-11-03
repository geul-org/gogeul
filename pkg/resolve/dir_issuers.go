//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.Issuers — the Issuers an archive holds documents of, with their names from the Issuer list
package resolve

import (
	"sort"
)

// Issuers lists the Issuer directories under `.well-known/geul/` (canonical
// IDs only) in ID order, each with the names the Issuer list (`issuers/`) binds to it.
func (d Dir) Issuers() ([]Named, error) {
	ents, err := d.store().ReadDir(".")
	if err != nil {
		return nil, err
	}
	b, _ := d.latest(issuersKey())
	names, _ := ParseIssuerNames(b)
	var out []Named
	for _, e := range ents {
		if e.IsDir() && isIssuerID(e.Name()) {
			out = append(out, Named{ID: e.Name(), Names: namesFor(names, e.Name())})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return len(out[i].ID) < len(out[j].ID) || len(out[i].ID) == len(out[j].ID) && out[i].ID < out[j].ID
	})
	return out, nil
}
