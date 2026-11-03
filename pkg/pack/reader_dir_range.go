//ff:func feature=pack type=resolver control=selection
//ff:what Reader.dirRange — the keys under a directory name without touching the index: <I> and <I>/docs (all of the Issuer's P keys), <I>/docs/<D> (one document), <I>/hashes (H keys)
package pack

import (
	"encoding/binary"
	"math"
	"strings"
)

// dirRange maps name to its keys without touching the index. `<I>` and
// `<I>/docs` hold the Issuer's P keys, `<I>/docs/<D>` one document's, and
// `<I>/hashes` the H keys (keyed by hash only, so a set of several Issuers
// answers for all). A name of no such shape, or under an Issuer outside the
// manifest's table, has no keys (it can still hold NAMED paths, hasUnder).
// The last document ID's range ends at the next Issuer, since its ID + 1 would
// wrap.
func (r *Reader) dirRange(name string) (keyRange, bool) {
	seg := strings.Split(name, "/")
	is, _ := normalIssuer(seg[0])
	num, known := r.inum[is]
	pathOf := func(k string) string {
		_, d, s := splitPKey(k)
		return pPath(is, d, s)
	}
	all := keyRange{r.m.pLayers, pKey(num, 0, 0), pKey(num+1, 0, 0), pathOf}
	switch {
	case !known:
		return keyRange{}, false
	case len(seg) == 1 || len(seg) == 2 && seg[1] == docsDir:
		return all, true
	case len(seg) == 2 && seg[1] == hashesDir:
		return keyRange{layers: r.m.hLayers, path: func(k string) string { return hPath(is, k) }}, true
	case len(seg) != 3 || seg[1] != docsDir || !isHex(seg[2], 16):
		return keyRange{}, false
	}
	id, _ := parseHex(seg[2], 16)
	d := binary.BigEndian.Uint64(id)
	if d != math.MaxUint64 {
		all.hi = pKey(num, d+1, 0)
	}
	all.lo = pKey(num, d, 0)
	return all, true
}
