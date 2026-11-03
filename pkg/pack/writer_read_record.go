//ff:func feature=pack type=writer control=sequence
//ff:what writer.readRecord — re-read a put, check it did not change since planning, and encode it (safe to run in parallel)
package pack

import (
	"crypto/sha256"
	"fmt"
)

// readRecord returns p's stored form.
func (w *writer) readRecord(p *put) (encoded, error) {
	data, err := w.b.Get(p.src)
	if err == nil && sha256.Sum256(data) != p.sum {
		err = fmt.Errorf("geul/pack: %s changed while writing", p.path)
	}
	if err != nil {
		return encoded{}, err
	}
	enc, b, err := w.encodeRecord(data)
	return encoded{enc: enc, b: b}, err
}
