//ff:type feature=pack type=model
//ff:what Dict — a zstd dictionary new records are compressed with: its ID, bytes and the encoders made for it
package pack

import "github.com/klauspost/compress/zstd"

// Dict is a dictionary (formatted, or raw content) under a DICTS ID. enc is
// the Zstd encoder made for it on first use, used the Zstd generation it was
// made at (Zstd.SetWorkers).
type Dict struct {
	ID    uint32
	Bytes []byte
	enc   *zstd.Encoder
	used  int
}
