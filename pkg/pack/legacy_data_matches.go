//ff:func feature=pack type=writer control=sequence
//ff:what legacyDataMatches — whether a sealed data pack begins with the open pack of a checkpoint without hash: its inventory starts with the checkpoint's entries and paths and goes on at the checkpoint's size
package pack

import (
	"io"
	"slices"
)

// legacyDataMatches checks a sealed data pack against a checkpoint that
// saved no hash: its header is a data pack's holding at least open.Count
// items, and its inventory (unframed without dictionary) begins with inv and
// paths, the next item (or the inventory) starting at open.Off.
func legacyDataMatches(f io.ReaderAt, open streamOpen, inv []streamInv, paths []string) bool {
	hb := make([]byte, headerSize)
	if _, err := f.ReadAt(hb, 0); err != nil {
		return false
	}
	h, err := parseObjectHeader(hb)
	if err != nil || h.kind != kindData || h.count < open.Count || len(inv) == 0 {
		return false
	}
	frame := make([]byte, h.invLen)
	if _, err := f.ReadAt(frame, int64(h.invOff)); err != nil {
		return false
	}
	raw, err := plainDecoder().DecodeAll(frame, nil)
	var got inventory
	if err == nil {
		got, err = parseInventory(raw)
	}
	if err != nil || len(got.entries) < len(inv) || len(got.paths) < len(paths) || !slices.Equal(got.paths[:len(paths)], paths) {
		return false
	}
	next := h.invOff
	if len(got.entries) > len(inv) {
		next = got.entries[len(inv)].off
	}
	return next == open.Off && invPrefixMatches(got.entries, inv)
}
