//ff:func feature=pack type=decoder control=sequence
//ff:what pLeafAt — whether the frame at a head's location is a P leaf whose first key is the head's
package pack

import "io"

// pLeafAt reads the block frame h.Loc names from f, unframes it (no
// dictionary) and checks it is a P leaf beginning with h.Key.
func pLeafAt(f io.ReaderAt, h streamKV) bool {
	frame := make([]byte, h.Loc.Len)
	if _, err := f.ReadAt(frame, int64(h.Loc.Off)); err != nil {
		return false
	}
	raw, err := plainDecoder().DecodeAll(frame, nil)
	if err != nil {
		return false
	}
	b, err := decodeBlock(raw)
	return err == nil && b.kind == blkPLeaf && len(b.keys) > 0 && b.keys[0] == h.Key
}
