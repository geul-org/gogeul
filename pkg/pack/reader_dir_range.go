//ff:func feature=pack type=resolver control=selection
//ff:what Reader.dirRange — the key range of a directory by its path's shape: `h`, `<Issuer>` or `<Issuer>/<문서>`; false for any other name
package pack

import (
	"encoding/binary"
	"math"
	"strings"
)

// dirRange maps name to its keys without touching the index. A name of no
// such shape, or under an Issuer outside the manifest's table, has no keys
// (it can still hold NAMED paths, hasUnder). The last document ID's range ends
// at the next Issuer, since its ID + 1 would wrap.
func (r *Reader) dirRange(name string) (keyRange, bool) {
	is, doc, nested := strings.Cut(name, "/")
	num, known := r.inum[is]
	pathOf := func(k string) string {
		_, d, s := splitPKey(k)
		return pPath(is, d, s)
	}
	end := pKey(num+1, 0, 0)
	id, canonical := parseHex(doc, 16)
	switch {
	case name == "h":
		return keyRange{layers: r.m.hLayers, path: hPath}, true
	case !known:
		return keyRange{}, false
	case !nested:
		return keyRange{r.m.pLayers, pKey(num, 0, 0), end, pathOf}, true
	case !canonical:
		return keyRange{}, false
	}
	d := binary.BigEndian.Uint64(id)
	if d != math.MaxUint64 {
		end = pKey(num, d+1, 0)
	}
	return keyRange{r.m.pLayers, pKey(num, d, 0), end, pathOf}, true
}
