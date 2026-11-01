//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what legacyIndexMatches — whether a file holds, back to back from the header, the P leaf blocks a checkpoint without hash wrote to its open index object
package pack

import "io"

// legacyIndexMatches checks an index object against a checkpoint that saved
// no hash: while records are put only P leaves go to index objects, so the
// open object's bytes up to open.Off are exactly the leaves whose heads name
// it, in order and back to back. Each must decode as a P leaf starting with
// its head's key (pLeafAt), and they must count open.Count and end at open.Off.
func legacyIndexMatches(f io.ReaderAt, open streamOpen, heads []streamKV) bool {
	off, n := uint64(headerSize), uint64(0)
	for _, h := range heads {
		if h.Loc.Obj != uint32(open.Num) {
			continue
		}
		if h.Loc.Off != off || !pLeafAt(f, h) {
			return false
		}
		off, n = off+h.Loc.Len, n+1
	}
	return off == open.Off && n == open.Count
}
