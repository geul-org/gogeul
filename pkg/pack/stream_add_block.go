//ff:func feature=pack type=writer control=sequence
//ff:what Stream.addBlock — write one directory block (zstd framed) to the open index object, starting a new one past PackSize
package pack

import "crypto/sha256"

// addBlock frames b with the codec (checksummed, §4.7) and appends it to the
// open index object, closing that object first when it holds blocks and
// would pass PackSize. It returns the block's location. The frame also goes
// into the object's running hash (isum).
func (s *Stream) addBlock(b block) (Loc, error) {
	frame, err := s.o.Build.Codec.Compress(encodeBlock(b), nil)
	if err != nil {
		return Loc{}, err
	}
	if s.st.Index.Num >= 0 && s.st.Index.Count > 0 && s.st.Index.Off+uint64(len(frame)) > uint64(s.o.Build.PackSize) {
		if err := s.closeObject(s.index, s.st.Index, kindIndex, nil); err != nil {
			return Loc{}, err
		}
		s.index, s.st.Index, s.isum = nil, streamOpen{Num: -1}, nil
	}
	if s.st.Index.Num < 0 {
		f, open, err := s.startObject(kindIndex, streamIndexPart)
		if err != nil {
			return Loc{}, err
		}
		s.index, s.st.Index, s.isum = f, open, sha256.New()
	}
	off := s.st.Index.Off
	if _, err := s.index.WriteAt(frame, int64(off)); err != nil {
		return Loc{}, err
	}
	s.isum.Write(frame)
	s.st.Index.Off += uint64(len(frame))
	s.st.Index.Count++
	return Loc{Obj: uint32(s.st.Index.Num), Off: off, Len: uint64(len(frame)), Enc: encZstd}, nil
}
