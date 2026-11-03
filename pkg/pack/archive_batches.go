//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Archive.batches — staged puts routed to their sets, staged removals sent to every set that has the path
package pack

import (
	"context"
	"slices"
)

// batches groups the staged changes per set; the caller holds a.mu.
func (a *Archive) batches(ctx context.Context) (map[*packSet]*Batch, error) {
	names := make([]string, 0, len(a.puts))
	for n := range a.puts {
		names = append(names, n)
	}
	slices.Sort(names)
	out := map[*packSet]*Batch{}
	add := func(s *packSet) *Batch {
		if out[s] == nil {
			out[s] = &Batch{Get: a.staged}
		}
		return out[s]
	}
	for _, n := range names {
		s, err := a.target(ctx, n)
		if err != nil {
			return nil, err
		}
		add(s).Put = append(add(s).Put, n)
	}
	for n := range a.dels {
		for _, s := range a.order(n) {
			add(s).Delete = append(add(s).Delete, n)
		}
	}
	return out, nil
}
