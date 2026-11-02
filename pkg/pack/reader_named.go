//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.named — a named path's record from the NAMED table (binary search)
package pack

import (
	"context"
	"sort"
)

// named reads a path outside the P and H spaces.
func (r *Reader) named(ctx context.Context, path string) ([]byte, error) {
	ns := r.m.named
	i := sort.Search(len(ns), func(i int) bool { return ns[i].path >= path })
	if i == len(ns) || ns[i].path != path {
		return nil, notFound(path)
	}
	return r.Record(ctx, ns[i].loc)
}
