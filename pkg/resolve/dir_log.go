//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.Log — a document's version history: every published version with its state, hash and publication time, and NEXT-LOCAL-ID
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// Log reads the version list of document doc of issuer and returns its
// versions in order, each with the RECORDED-AT of the list that first carried
// its PUBLICATION (reached by PARENT through `<Issuer>/hashes/`), and the list's
// NEXT-LOCAL-ID. A document never published is ErrNotFound.
func (d Dir) Log(issuer, doc []core.Word) (History, error) {
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	lg, err := d.readList(k, k.Doc)
	if err != nil {
		return History{}, err
	}
	if !lg.Present {
		return History{}, ErrNotFound
	}
	b, _ := d.Fetch(context.Background(), listPath(k))
	ts := d.listTimes(k.Issuer, b, len(lg.Records)+1)
	h := History{Next: lg.Next}
	for j, r := range lg.Records {
		p, err := core.DecodePublication(r)
		if err != nil || p.Codec != core.EdgePublication {
			continue
		}
		e, _ := entryOf(lg.Entries, p.Version)
		h.Entries = append(h.Entries, LogEntry{VersionEntry: e, At: timeOf(ts, j)})
	}
	return h, nil
}
