//ff:func feature=pack type=resolver control=selection
//ff:what Reader.Has — whether this generation serves a path, without reading its record
package pack

import (
	"context"
	"errors"
	"sort"
)

// Has reports whether path is live in r (index only; no record is read).
func (r *Reader) Has(ctx context.Context, path string) (bool, error) {
	ref := parsePath(path)
	var err error
	switch ref.space {
	case spaceP:
		num, ok := r.inum[ref.issuer]
		if !ok {
			return false, nil
		}
		_, err = r.find(ctx, r.m.pLayers, pKey(num, ref.doc, ref.slot), path)
	case spaceH:
		_, err = r.find(ctx, r.m.hLayers, ref.hash, path)
	default:
		path = ref.path
		i := sort.Search(len(r.m.named), func(i int) bool { return r.m.named[i].path >= path })
		return i < len(r.m.named) && r.m.named[i].path == path, nil
	}
	if errors.Is(err, ErrNotFound) {
		return false, nil
	}
	return err == nil, err
}
