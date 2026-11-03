//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what canonSegs — every segment of a split path in its stored spelling (canonSeg)
package pack

// canonSegs rewrites each segment of seg with canonSeg, in place, and returns
// seg.
func canonSegs(seg []string) []string {
	for i, s := range seg {
		seg[i] = canonSeg(s)
	}
	return seg
}
