//ff:func feature=pack type=helper control=sequence
//ff:what Zstd.workers — how many records compress at once: Workers when set, else DefaultWorkers but no more than GOMAXPROCS
package pack

import "runtime"

// workers is z.Workers when at least 1. Otherwise it is DefaultWorkers, or
// GOMAXPROCS when that is smaller: more encoders than threads only hold
// memory (each best-level encoder holds about 250 MB with a dictionary).
// A writer that wants more encoders sets Workers (or SetWorkers); the frames
// are the same whatever the count.
func (z *Zstd) workers() int {
	return defaultInt(z.Workers, 1, min(DefaultWorkers, runtime.GOMAXPROCS(0)))
}
