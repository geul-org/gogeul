//ff:func feature=pack type=encoder control=sequence
//ff:what crashCodec.Compress — the inner codec's frame while calls are left, then an error
package pack

import "errors"

// Compress passes src to the inner codec until the calls run out.
func (c crashCodec) Compress(src []byte, d *Dict) ([]byte, error) {
	if *c.left <= 0 {
		return nil, errors.New("crash")
	}
	*c.left--
	return c.inner.Compress(src, d)
}
