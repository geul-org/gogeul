//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.anyLive — whether kr has a live key outside gone, stopping at the first one and keeping no key
package pack

import (
	"context"
	"errors"
)

// anyLive walks kr's layers in key order until a key passes checkLive. Nothing
// is collected, so memory is the block cache's however large the range is.
func (r *Reader) anyLive(ctx context.Context, kr keyRange, gone map[string]bool) (bool, error) {
	check := func(k string, l Loc) error { return r.checkLive(ctx, kr, gone, k, l) }
	for _, l := range kr.layers {
		err := r.walk(ctx, l.blk, kr.lo, kr.hi, 0, check)
		if errors.Is(err, errLiveKey) {
			return true, nil
		}
		if err != nil {
			return false, err
		}
	}
	return false, nil
}
