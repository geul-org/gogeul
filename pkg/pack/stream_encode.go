//ff:func feature=pack type=writer control=sequence
//ff:what Stream.Encode — a record ready to Put: zstd with the stream's dictionary when smaller, else raw, with its GEUL content hash; safe from many goroutines
package pack

// Encode prepares b for Put or PutNamed: compressed with the stream's codec
// and dictionary when that is smaller (Stored keeps it raw), and its GEUL
// content hash when b is a document (contentHash), under which the record
// also serves h/. It may run on many goroutines at once while one goroutine
// puts the results in key order. Once Close has begun the record is not
// stored again (closedPut), so it is left raw.
func (s *Stream) Encode(b []byte) (StreamRecord, error) {
	r := StreamRecord{Enc: encRaw, Bytes: b, Raw: len(b)}
	r.Hash, r.HashOK = contentHash(b)
	if _, ok := s.o.Build.Codec.(Stored); ok || s.st.Close.Phase > closeNone {
		return r, nil
	}
	z, err := s.o.Build.Codec.Compress(b, s.dict)
	if err != nil {
		return StreamRecord{}, err
	}
	if len(z) < len(b) {
		r.Enc, r.Bytes = encZstd, z
	}
	return r, nil
}
