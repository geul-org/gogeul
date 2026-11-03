//ff:func feature=pack type=helper control=sequence
//ff:what TestZstdWorkers — 0 is DefaultWorkers capped at GOMAXPROCS; a set value stays, above both caps too
package pack

import (
	"runtime"
	"testing"
)

func TestZstdWorkers(t *testing.T) {
	old := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(old)
	if got := (&Zstd{}).workers(); got != 1 {
		t.Fatalf("GOMAXPROCS 1: %d", got)
	}
	runtime.GOMAXPROCS(DefaultWorkers + 4)
	if got := (&Zstd{}).workers(); got != DefaultWorkers {
		t.Fatalf("GOMAXPROCS %d: %d", DefaultWorkers+4, got)
	}
	if (&Zstd{Workers: 3}).workers() != 3 || (&Zstd{Workers: 64}).workers() != 64 || (&Zstd{Workers: -1}).workers() != DefaultWorkers {
		t.Fatal("set workers")
	}
}
