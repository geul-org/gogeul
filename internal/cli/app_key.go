//ff:func feature=cli type=helper control=sequence
//ff:what app.key — a document address to its Key: a local Issuer key name first, then the archive's IDs and names
package cli

import "github.com/geul-org/gogeul/pkg/resolve"

// key resolves `<Issuer>/<문서>[@<버전>]`. An Issuer part that names a local
// key becomes that Issuer's ID (localIssuer); the rest is the data folder's
// Lookup (IDs only without a data folder).
func (a *app) key(s string) (resolve.Key, error) {
	s, err := a.localIssuer(s)
	if err != nil {
		return resolve.Key{}, err
	}
	d, _ := a.dir()
	return d.Lookup(s)
}
