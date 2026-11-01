//ff:func feature=pack type=writer control=sequence
//ff:what streamCrashStart — a crash test stream in out: documents 1–25 and a checkpoint "25", the checkpoint made hash-less (as gogeul 6c02e78 wrote it) when legacy
package pack

import "testing"

// streamCrashStart starts the stream with streamCrashOptions, puts documents
// 1–25 and checkpoints. With legacy the saved state loses its hash states
// (streamTestLegacy) and its bytes are returned, to put back after a Close
// that the old code would not have saved.
func streamCrashStart(t *testing.T, out string, legacy bool) (*Stream, []byte) {
	t.Helper()
	s, err := NewStream(out, streamCrashOptions())
	if err != nil {
		t.Fatal(err)
	}
	streamTestPut(t, s, 1, 25)
	if err := s.Checkpoint([]byte("25")); err != nil {
		t.Fatal(err)
	}
	if !legacy {
		return s, nil
	}
	return s, streamTestLegacy(t, out)
}
