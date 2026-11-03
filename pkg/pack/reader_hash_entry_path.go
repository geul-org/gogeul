//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.hashEntryPath — the path of an H key: under the set's one Issuer without reading, else under the Issuer its bytes name
package pack

import "context"

// hashEntryPath names the H key k at l. A set of one Issuer names it without
// reading; in a set of several the record is read (HashPath), so the index
// and the inventories name it alike.
func (r *Reader) hashEntryPath(ctx context.Context, k string, l Loc) (string, error) {
	if len(r.m.issuers) == 1 {
		return hPath(r.m.issuers[0], k), nil
	}
	return r.HashPath(ctx, l)
}
