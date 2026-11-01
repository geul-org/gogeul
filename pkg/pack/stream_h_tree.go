//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.hTree — the H root over the leaf heads in the head file: the first internal level written group by group from the file, the rest by tree
package pack

import "io"

// hTree builds the H internal levels over count leaf heads in f, as tree
// does over a slice: when the heads fit the root (HRoot 0, at most HRoot, or
// at most HFanout) tree gets them all; otherwise the first level is written
// HFanout heads at a time read from f, and tree goes on over that level,
// which holds one head per HFanout leaves. So at most max(HRoot, HFanout)
// heads are in memory (all of them only with HRoot 0, a root over all).
func (s *Stream) hTree(f io.ReaderAt, count int64) (block, error) {
	fanout, rootMax := s.o.Build.HFanout, s.o.Build.HRoot
	if rootMax == 0 || count <= int64(rootMax) || count <= int64(fanout) {
		heads, err := readHeads(f, 0, count)
		if err != nil {
			return block{}, err
		}
		return s.tree(heads, blkHInner, fanout, rootMax)
	}
	var next []streamKV
	for i := int64(0); i < count; i += int64(fanout) {
		g, err := readHeads(f, i, min(int64(fanout), count-i))
		if err != nil {
			return block{}, err
		}
		loc, err := s.addBlock(blockOf(blkHInner, g))
		if err != nil {
			return block{}, err
		}
		next = append(next, streamKV{Key: g[0].Key, Loc: loc})
	}
	return s.tree(next, blkHInner, fanout, rootMax)
}
