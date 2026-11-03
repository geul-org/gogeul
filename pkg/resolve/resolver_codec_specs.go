//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Resolver.CodecSpecs — the codec specs of the given targets, each found under the GEUL Issuer 0x00000000 by its TAG through the source list
package resolve

import (
	"context"
	"errors"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// CodecSpecs fetches the newest spec document of each target by its address
// (codecDocID, GEUL-ORG 규격 문서의 위치) through the source list, so a spec
// missing from the first source comes from the next. A target no source has
// is left out, so its records read generically. Any other error stops it.
func (r *Resolver) CodecSpecs(ctx context.Context, targets []core.Word) (codecspec.Spec, error) {
	s := codecspec.Spec{}
	for _, t := range targets {
		n, err := r.Fetch(ctx, Key{Issuer: []core.Word{0, 0}, Doc: codecDocID(t)})
		if errors.Is(err, ErrNotFound) {
			continue
		}
		if err != nil {
			return nil, err
		}
		s = s.Merge(specOfDoc(n.Doc))
	}
	return s, nil
}
