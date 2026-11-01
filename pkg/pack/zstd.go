//ff:type feature=pack type=model
//ff:what Zstd — the compressing Codec (klauspost/compress): level, dictionary size, encoders made on first use
package pack

import (
	"sync"

	"github.com/klauspost/compress/zstd"
)

// Zstd compresses with klauspost/compress. Level is a zstd level 1–22 (0 is
// DefaultZstdLevel); the library maps it to its four speeds (1–2 fastest,
// 3–6 default, 7–10 better, 11+ best). DictSize is the size of a dictionary
// trained for a set that has none (0 is DefaultDictSize); below 0 new records
// use no dictionary. Workers is how many records a writer compresses at once
// (0 is DefaultWorkers; each best-level encoder holds tens of MiB); SetWorkers
// changes it while records compress. gen counts those changes and used is the
// gen plain was made at. Use it as *Zstd; it is safe for concurrent use.
type Zstd struct {
	Level    int
	DictSize int
	Workers  int

	mu    sync.Mutex
	plain *zstd.Encoder
	used  int
	gen   int
}
