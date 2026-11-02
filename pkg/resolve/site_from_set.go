//ff:func feature=resolve type=resolver control=sequence
//ff:what site.fromSet — a path from one named set; ErrNotFound when that set does not exist or lacks it
package resolve

import "context"

// fromSet reads path from the set name.
func (s *site) fromSet(ctx context.Context, name, path string) ([]byte, error) {
	r, err := s.set(ctx, name)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return nil, ErrNotFound
	}
	return r.Fetch(ctx, path)
}
