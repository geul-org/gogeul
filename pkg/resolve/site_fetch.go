//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what site.Fetch — Source: a path from the first pack set that has it, then as a loose file; ErrNotFound only when neither does
package resolve

import (
	"context"
	"errors"
)

// Fetch tries the sets of path in order, then the loose file. A set that is
// there but broken stops it with that error, never ErrNotFound: a source that
// cannot be read is not a source that lacks the document.
func (s *site) Fetch(ctx context.Context, path string) ([]byte, error) {
	names, err := s.setsFor(ctx, path)
	if err != nil {
		return nil, err
	}
	for _, n := range names {
		b, err := s.fromSet(ctx, n, path)
		if err == nil {
			return b, nil
		}
		if !errors.Is(err, ErrNotFound) {
			return nil, err
		}
	}
	return s.loose.Fetch(ctx, path)
}
