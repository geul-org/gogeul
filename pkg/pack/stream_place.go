//ff:func feature=pack type=writer control=sequence
//ff:what Stream.place — append stored bytes to the open data pack, closing it and starting the next when they would pass PackSize
package pack

import "crypto/sha256"

// place writes b (stored with encoding enc) to the open data pack and returns
// its location. A pack that already holds items and would pass PackSize is
// finished first (§6 step 4); an item alone may be bigger. The bytes also
// go into the pack's running hash (dsum).
func (s *Stream) place(b []byte, enc byte) (Loc, error) {
	if s.st.Data.Num >= 0 && s.st.Data.Count > 0 && s.st.Data.Off+uint64(len(b)) > uint64(s.o.Build.PackSize) {
		if err := s.finishData(); err != nil {
			return Loc{}, err
		}
	}
	if s.st.Data.Num < 0 {
		f, open, err := s.startObject(kindData, streamDataPart)
		if err != nil {
			return Loc{}, err
		}
		s.data, s.st.Data, s.dsum = f, open, sha256.New()
	}
	off := s.st.Data.Off
	if _, err := s.data.WriteAt(b, int64(off)); err != nil {
		return Loc{}, err
	}
	s.dsum.Write(b)
	s.st.Data.Off += uint64(len(b))
	s.st.Data.Count++
	return Loc{Obj: uint32(s.st.Data.Num), Off: off, Len: uint64(len(b)), Enc: enc}, nil
}
