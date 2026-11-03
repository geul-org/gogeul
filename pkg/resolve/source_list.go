//ff:func feature=resolve type=resolver control=sequence
//ff:what sourceList — a document's versions from the version list document of one source
package resolve

import "context"

// sourceList fetches d's version list document from e and returns its versions.
// A list that does not read is a verification failure.
func sourceList(ctx context.Context, e Entry, d DocKey) ([]VersionEntry, error) {
	lb, err := e.Source.Fetch(ctx, listPath(d))
	if err != nil {
		return nil, err
	}
	lg, err := decodeVersionList(lb, d.Issuer, d.Doc)
	if err != nil {
		return nil, err
	}
	return lg.Entries, nil
}
