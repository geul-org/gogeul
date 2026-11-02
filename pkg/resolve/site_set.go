//ff:func feature=resolve type=resolver control=selection
//ff:what site.set — the Reader of a set name ("" is the top set), opened on first use; nil when no latest.txt is there
package resolve

import (
	"context"
	"errors"

	"github.com/geul-org/gogeul/pkg/pack"
)

// set opens name once. A missing latest.txt is remembered as no set; a broken
// set is not remembered, so its error stays and a passing failure can recover.
func (s *site) set(ctx context.Context, name string) (*pack.Reader, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if r, ok := s.sets[name]; ok {
		return r, nil
	}
	r, err := pack.Open(ctx, subStore{base: s.objects, name: name}, pack.Options{})
	switch {
	case errors.Is(err, ErrNotFound):
		s.sets[name] = nil
		return nil, nil
	case err != nil:
		return nil, err
	}
	s.sets[name] = r
	return r, nil
}
