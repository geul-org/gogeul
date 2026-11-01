//ff:func feature=pack type=writer control=sequence
//ff:what Stream.closeObject — append the tail, write the header, sync, name the part file by its SHA-256 under o/ and fill its OBJECTS entry
package pack

import (
	"os"
	"path/filepath"
)

// closeObject finishes the open object (f, open) of kind: tail (a data
// pack's inventory, nil for an index object) after its items, the header
// (§4.4), then the file is synced and renamed to o/<sha256>.<ext> (a resume
// may restore an open object from it, reopenPart) and its reserved
// OBJECTS entry gets size and hash.
func (s *Stream) closeObject(f *os.File, open streamOpen, kind byte, tail []byte) error {
	h := objectHeader{kind: kind, count: open.Count}
	if kind == kindData {
		h.invOff, h.invLen = open.Off, uint64(len(tail))
	}
	h.size = open.Off + uint64(len(tail))
	_, err := f.WriteAt(tail, int64(open.Off))
	if err == nil {
		_, err = f.WriteAt(encodeObjectHeader(h), 0)
	}
	if err == nil {
		err = f.Truncate(int64(h.size))
	}
	var sum [32]byte
	if err == nil {
		sum, err = sumFile(f, h.size)
	}
	if err == nil {
		err = f.Sync()
	}
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	o := object{kind: kind, size: h.size, hash: sum}
	if err == nil {
		err = os.Rename(f.Name(), filepath.Join(s.out, filepath.FromSlash(objectName(o))))
	}
	if err != nil {
		return err
	}
	s.st.Objects[open.Num] = streamObject{Kind: kind, Size: h.size, Hash: sum}
	s.st.Res.Objects = append(s.st.Res.Objects, objectName(o))
	s.st.Res.Bytes += int64(h.size)
	return nil
}
