//ff:func feature=pack type=writer control=sequence
//ff:what writer.planPut — read one Put path; nil when its bytes equal the previous generation's
package pack

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// planPut returns the put for path or nil when unchanged.
func (w *writer) planPut(path string) (*put, error) {
	if path == "" || len(path) > 0xFFFF {
		return nil, fmt.Errorf("geul/pack: invalid path %q", path)
	}
	data, err := w.b.Get(path)
	if err != nil {
		return nil, err
	}
	w.sums[path] = sha256.Sum256(data)
	if l, ok := w.base[path]; ok {
		old, err := w.prev.Record(w.ctx, l)
		if err != nil {
			return nil, err
		}
		if bytes.Equal(old, data) {
			return nil, nil
		}
	}
	pu := &put{path: path, ref: parsePath(path), sum: w.sums[path], n: len(data)}
	pu.ch, pu.chOK = contentHash(data)
	return pu, nil
}
