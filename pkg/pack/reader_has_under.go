//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.hasUnder — whether a live path outside gone lies under directory name: a NAMED path, else a key of its range
package pack

import (
	"context"
	"sort"
	"strings"
)

// hasUnder answers from the manifest's NAMED table, then from the keys the
// path's shape gives (dirRange); it never lists the generation's paths.
func (r *Reader) hasUnder(ctx context.Context, name string, gone map[string]bool) (bool, error) {
	prefix := dirPrefix(name)
	ns := r.m.named
	i := sort.Search(len(ns), func(i int) bool { return ns[i].path >= prefix })
	for ; i < len(ns) && strings.HasPrefix(ns[i].path, prefix); i++ {
		if !gone[ns[i].path] {
			return true, nil
		}
	}
	kr, ok := r.dirRange(name)
	if !ok {
		return false, nil
	}
	return r.anyLive(ctx, kr, gone)
}
