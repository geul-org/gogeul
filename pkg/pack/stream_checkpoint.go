//ff:func feature=pack type=writer control=sequence
//ff:what Stream.Checkpoint — spill the held H keys, flush the open objects to disk and save the state with the caller's progress, so ResumeStream continues from here
package pack

// Checkpoint makes everything put so far durable: the held H keys become a
// run, the open data pack and index object are synced and the state
// (streamState) is saved atomically. After a crash ResumeStream continues
// from the last checkpoint; records put after it are put again by the caller.
// meta is the caller's progress (what it has put, its own files' sizes),
// saved in the same atomic write so the two never disagree (Stream.Meta).
// Checkpoints do not change the bytes the stream writes. Once Close has
// begun only meta is saved: the records are all in already.
func (s *Stream) Checkpoint(meta []byte) error {
	s.st.Meta = meta
	if s.st.Close.Phase > closeNone {
		return s.saveState()
	}
	if err := s.spill(); err != nil {
		return err
	}
	if err := syncFile(s.data); err != nil {
		return err
	}
	if err := syncFile(s.index); err != nil {
		return err
	}
	return s.saveState()
}
