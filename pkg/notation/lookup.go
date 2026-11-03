//ff:type feature=notation type=model
//ff:what Lookup — how a document address (Issuer ID or name, document ID, optional @version) becomes a Key when reading tokens
package notation

import "github.com/geul-org/gogeul/pkg/resolve"

// Lookup resolves `<Issuer>/<문서>[@<버전>]` (resolve.Dir.Lookup). With a
// nil Lookup only canonical IDs are read.
type Lookup func(addr string) (resolve.Key, error)
