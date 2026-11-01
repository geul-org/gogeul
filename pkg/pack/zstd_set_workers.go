//ff:func feature=pack type=encoder control=sequence
//ff:what Zstd.SetWorkers — change how many records compress at once: the kept encoders go, and are made again on next use for that many
package pack

// SetWorkers makes Compress hold n encoders from now on (below 1: no change).
// An encoder holds its memory once it has been used and they are used in
// turn, so fewer records at a time only frees memory when there are fewer
// encoders. The old ones are dropped once the calls in progress finish.
// EncodeAll gives the same frame whatever the encoder count, so the bytes do
// not change.
func (z *Zstd) SetWorkers(n int) {
	if n < 1 {
		return
	}
	z.mu.Lock()
	defer z.mu.Unlock()
	z.Workers, z.plain = n, nil
	z.gen++
}
