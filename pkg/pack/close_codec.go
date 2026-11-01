//ff:func feature=pack type=writer control=sequence
//ff:what closeCodec — the codec Close compresses blocks with: for Zstd a one-encoder copy at the same level, the original's cached encoders dropped; others as they are
package pack

// closeCodec returns the codec for Close. Zstd keeps up to Workers encoders
// for the stream's records, the dictionary's among them (each best-level
// one holds tens to hundreds of MiB); Close compresses one block at a time
// and no record, so those are dropped (made again on next use) and a Zstd
// with one encoder at the same level is used. EncodeAll gives the same
// frame whatever the encoder count, so the bytes do not change.
func closeCodec(c Codec, d *Dict) Codec {
	z, ok := c.(*Zstd)
	if !ok {
		return c
	}
	z.mu.Lock()
	z.plain = nil
	if d != nil {
		d.enc = nil
	}
	z.mu.Unlock()
	return &Zstd{Level: z.Level, DictSize: z.DictSize, Workers: 1}
}
