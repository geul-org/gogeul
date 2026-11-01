//ff:func feature=pack type=helper control=sequence
//ff:what peakHeap — the most heap bytes in objects seen while f runs (collector kept tight), above what was there before it
package pack

import (
	"runtime"
	"runtime/debug"
	"runtime/metrics"
	"time"
)

// peakHeap collects garbage, reads the heap object bytes, then samples them
// every 100µs while f runs and returns the highest sample less the first.
// GOGC is 10 meanwhile, so the samples follow live memory, not the
// collector's slack.
func peakHeap(f func()) uint64 {
	defer debug.SetGCPercent(debug.SetGCPercent(10))
	runtime.GC()
	s := []metrics.Sample{{Name: "/memory/classes/heap/objects:bytes"}}
	metrics.Read(s)
	base, peak := s[0].Value.Uint64(), uint64(0)
	done := make(chan struct{})
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		for {
			metrics.Read(s)
			peak = max(peak, s[0].Value.Uint64())
			select {
			case <-done:
				return
			case <-time.After(100 * time.Microsecond):
			}
		}
	}()
	f()
	close(done)
	<-stopped
	return max(peak, base) - base
}
