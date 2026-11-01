//ff:func feature=pack type=writer control=sequence
//ff:what Stream.closeBegin — drop the record encoders and use one block encoder, then (first time only) spill the held H keys, sync the open objects and save the state as Close's starting point
package pack

import "runtime/debug"

// closeBegin readies Close. Records are no longer compressed, so the
// codec's cached encoders are dropped and blocks are compressed by one
// encoder (closeCodec: the same bytes), and memory freed by the caller's
// bulk work is returned to the OS. A Close that runs for the first time then
// spills the held H keys, syncs the open objects and saves the state with
// Phase closeStarted: from here a resume keeps every record (closedPut).
func (s *Stream) closeBegin() error {
	s.o.Build.Codec = closeCodec(s.o.Build.Codec, s.dict)
	debug.FreeOSMemory()
	if s.st.Close.Phase > closeNone {
		return nil
	}
	if err := s.spill(); err != nil {
		return err
	}
	s.st.Close.Phase = closeStarted
	return s.closeSave()
}
