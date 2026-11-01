//ff:func feature=pack type=encoder control=sequence
//ff:what Zstd.encoder — the encoder for no dictionary or for d, made on first use and kept until SetWorkers changes the count
package pack

import "github.com/klauspost/compress/zstd"

// encoder returns z's plain encoder (d nil) or d's, made again when it was
// made before the last SetWorkers.
func (z *Zstd) encoder(d *Dict) (*zstd.Encoder, error) {
	z.mu.Lock()
	defer z.mu.Unlock()
	slot, used := &z.plain, &z.used
	if d != nil {
		slot, used = &d.enc, &d.used
	}
	if *slot == nil || *used != z.gen {
		e, err := newEncoder(z.level(), z.workers(), d)
		if err != nil {
			return nil, err
		}
		*slot, *used = e, z.gen
	}
	return *slot, nil
}
