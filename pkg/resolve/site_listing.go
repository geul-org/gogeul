//ff:func feature=resolve type=resolver control=sequence
//ff:what site.listing — what lists a local site: an Archive of every set over the loose files; a remote site cannot be listed
package resolve

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"
)

// listing opens every set under the folder. Only listings need that; reads
// open the sets they touch.
func (s *site) listing() (Store, error) {
	if s.dir == "" {
		return nil, fmt.Errorf("geul/resolve: a remote source has no listing")
	}
	a, err := pack.OpenDir(context.Background(), s.dir, pack.BuildOptions{})
	if err != nil {
		return nil, err
	}
	return overlay{stored: a, loose: dirStore{root: s.dir}}, nil
}
