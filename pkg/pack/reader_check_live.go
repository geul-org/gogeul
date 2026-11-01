//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.checkLive — errLiveKey when the visited key is live: no marker, not removed in gone, not hidden by a newer layer's marker
package pack

import (
	"context"
	"errors"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// checkLive judges one leaf entry of kr during anyLive's walk. gone holds
// paths staged for removal. Nil means walk on.
func (r *Reader) checkLive(ctx context.Context, kr keyRange, gone map[string]bool, k string, l Loc) error {
	if l.Enc == encGone || gone[kr.path(k)] {
		return nil
	}
	_, err := r.find(ctx, kr.layers, k, kr.path(k))
	if errors.Is(err, resolve.ErrNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	return errLiveKey
}
