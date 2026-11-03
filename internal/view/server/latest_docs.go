//ff:func feature=server type=helper control=sequence
//ff:what latestDocs — the newest version of an Issuer's document list in the tree (latest.geul through its versions.geul)
package server

import (
	"context"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// latestDocs reads the document list of the Issuer folder space (`<space>/docs/`)
// at its newest version. A space that is not an Issuer ID, or a list that is
// not there, is an error.
func latestDocs(t resolve.Store, space string) ([]byte, error) {
	a, err := resolve.ParseAddr(space + "/docs/latest.geul")
	if err != nil {
		return nil, err
	}
	return resolve.FetchAddr(context.Background(), resolve.Dir{Store: t}, a)
}
