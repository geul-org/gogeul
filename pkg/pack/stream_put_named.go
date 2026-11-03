//ff:func feature=pack type=writer control=sequence
//ff:what Stream.PutNamed — store a record at a named path (outside the P and H spaces), once per path
package pack

import "fmt"

// PutNamed stores r at the named path (§3.4: a list `issuers/<V>.geul`,
// `<Issuer>/docs/<V>.geul`, `<Issuer>/codecs/<V>.geul` or its versions.geul, or a
// codec `<Issuer>/codecs/<코덱 ID>/<V>.geul` or its versions.geul),
// in the open data pack's path table. A path may be put once; NAMED is
// sorted when the stream closes. A named record does not serve h/. Once
// Close has begun, a path it holds is acknowledged without storing
// (closedPutNamed).
func (s *Stream) PutNamed(path string, r StreamRecord) error {
	ref := parsePath(path)
	if ref.space != spaceNamed {
		return fmt.Errorf("geul/pack: %q is not a named path", path)
	}
	path = ref.path
	if s.st.Close.Phase > closeNone {
		return s.closedPutNamed(path)
	}
	loc, err := s.place(r.Bytes, r.Enc)
	if err != nil {
		return err
	}
	s.st.Inv = append(s.st.Inv, streamInv{Kind: invNamed, Enc: r.Enc, Ref: uint32(len(s.st.Paths)), N: uint32(len(r.Bytes)), Off: loc.Off})
	s.st.Paths = append(s.st.Paths, path)
	s.st.Named = append(s.st.Named, streamKV{Key: path, Loc: loc})
	r.HashOK = false
	s.count(r)
	return nil
}
