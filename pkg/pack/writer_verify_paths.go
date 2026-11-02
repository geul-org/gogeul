//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.verifyPaths — every Put path reads back with its SHA-256, every deletion is NotFound, a Full batch is the whole path set
package pack

import (
	"crypto/sha256"
	"errors"
)

// verifyPaths checks the path contract (§3.1) on the new generation r.
func (w *writer) verifyPaths(r *Reader) error {
	for p, sum := range w.sums {
		b, err := r.Fetch(w.ctx, p)
		if err != nil {
			return err
		}
		if sha256.Sum256(b) != sum {
			return corrupt("verify: bytes differ at " + p)
		}
	}
	for _, d := range w.dels {
		if _, err := r.Fetch(w.ctx, d); !errors.Is(err, ErrNotFound) {
			return corrupt("verify: deleted path still served: " + d)
		}
	}
	if !w.b.Full {
		return nil
	}
	paths, err := r.Paths(w.ctx)
	if err == nil && len(paths) != len(w.sums) {
		err = corrupt("verify: path set differs from the batch")
	}
	return err
}
