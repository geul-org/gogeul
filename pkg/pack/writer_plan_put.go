//ff:func feature=pack type=writer control=sequence
//ff:what writer.planPut — read one Put path; nil when its bytes equal the previous generation's
package pack

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// planPut returns the put for path or nil when unchanged. A path outside the
// 공표 주소 forms (§3.13), latest.geul included, is refused: it is never
// stored. A zero-padded Issuer is stored under its canonical spelling.
func (w *writer) planPut(path string) (*put, error) {
	ref := parsePath(path)
	if ref.space == spaceNone || len(path) > 0xFFFF {
		return nil, fmt.Errorf("geul/pack: %q is not a stored 공표 주소", path)
	}
	src := path
	data, err := w.b.Get(src)
	if err != nil {
		return nil, err
	}
	path = ref.path
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
	pu := &put{path: path, src: src, ref: ref, sum: w.sums[path], n: len(data)}
	pu.ch, pu.chOK = contentHash(data)
	return pu, nil
}
