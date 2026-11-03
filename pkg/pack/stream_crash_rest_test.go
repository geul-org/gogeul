//ff:func feature=pack type=writer control=sequence
//ff:what streamCrashRest — put documents 26–40 and the named record into a crash test stream
package pack

import "testing"

// streamCrashRest puts what follows the checkpoint: documents 26–40 and
// 0x00010000/docs/0x00000001.geul.
func streamCrashRest(t *testing.T, s *Stream) {
	t.Helper()
	streamTestPut(t, s, 26, 40)
	r, _ := s.Encode([]byte("q1 0x5100000000000001\n"))
	if err := s.PutNamed("0x00000001/docs/0x00000001.geul", r); err != nil {
		t.Fatal(err)
	}
}
