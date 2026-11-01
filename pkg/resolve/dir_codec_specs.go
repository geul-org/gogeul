//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.CodecSpecs — the codec specs of the given targets only, each read by its address without listing the GEUL Issuer
package resolve

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// CodecSpecs gathers the latest spec document of each target (codecKey), for
// a reader of a few documents; CodecSpec gathers every one by listing. A
// target with no live spec document is left out, so its records read
// generically. Any other error stops it.
func (d Dir) CodecSpecs(targets []core.Word) (codecspec.Spec, error) {
	s := codecspec.Spec{}
	for _, t := range targets {
		spec, err := d.codecSpecOf(t)
		if errors.Is(err, ErrNotFound) {
			continue
		}
		if err != nil {
			return nil, err
		}
		s = s.Merge(spec)
	}
	return s, nil
}
