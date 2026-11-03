//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what parseIssuers — an Issuer table: u32 count, entries of width (2·4·16 words) and big-endian ID bytes, canonical only
package pack

import "fmt"

// parseIssuers decodes an Issuer table into canonical hex; duplicates, other
// widths (2, 4 or 16 words only) and zero-padded entries are ErrCorrupt. Trailing bytes are left in c.
func parseIssuers(c *cursor) ([]string, error) {
	n := c.u32()
	if c.bad || uint64(n)*5 > uint64(len(c.b)) {
		return nil, corrupt("Issuer table")
	}
	out := make([]string, n)
	seen := map[string]bool{}
	for i := range out {
		w := c.u8()
		b := c.take(uint64(w) * 2)
		id := fmt.Sprintf("0x%X", b)
		if c.bad || w != 2 && w != 4 && w != 16 || seen[id] || !canonicalIssuer(id) {
			return nil, corrupt("Issuer table entry")
		}
		out[i], seen[id] = id, true
	}
	return out, nil
}
