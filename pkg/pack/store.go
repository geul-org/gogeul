//ff:type feature=pack type=resolver
//ff:what Store — where pack objects live: whole reads (latest.txt, manifest) and single-range reads (blocks, records)
package pack

import "context"

// Store reads objects by their name under a pack base (`latest.txt`,
// `m/<sha256>.gpm`, `o/<sha256>.gpk`). A missing object is an error wrapping
// ErrNotFound. ReadRange must return exactly n bytes from off or fail.
type Store interface {
	ReadAll(ctx context.Context, name string) ([]byte, error)
	ReadRange(ctx context.Context, name string, off, n uint64) ([]byte, error)
}
