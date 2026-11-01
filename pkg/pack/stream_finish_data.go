//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.finishData — close the open data pack with its inventory (one checksummed zstd frame) and start a fresh inventory
package pack

// finishData writes the open data pack's inventory (§4.5: the stream's
// Issuers, the named paths it holds, its entries in offset order) and closes
// the pack.
func (s *Stream) finishData() error {
	inv := inventory{issuers: s.st.Issuers, paths: s.st.Paths}
	for _, e := range s.st.Inv {
		inv.entries = append(inv.entries, invEntry{kind: e.Kind, enc: e.Enc, flags: e.Flags, ref: e.Ref, doc: e.Doc, slot: e.Slot, n: e.N, off: e.Off})
	}
	tail, err := s.o.Build.Codec.Compress(encodeInventory(inv), nil)
	if err == nil {
		err = s.closeObject(s.data, s.st.Data, kindData, tail)
	}
	if err != nil {
		return err
	}
	s.data, s.st.Data, s.st.Inv, s.st.Paths, s.dsum = nil, streamOpen{Num: -1}, nil, nil, nil
	return nil
}
