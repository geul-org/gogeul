//ff:func feature=pack type=writer control=sequence
//ff:what streamTestHPass — a test stream at the start of an H pass over runs of the given keys, with the pass opened
package pack

import "testing"

// streamTestHPass writes the runs (streamTestRuns), starts the pass
// (hStart) and opens it (openHPass).
func streamTestHPass(t *testing.T, c Codec, keys ...[]byte) (*Stream, *hPass) {
	t.Helper()
	s := streamTestNew(t, c)
	s.st.Runs = streamTestRuns(t, s, keys...)
	if err := s.hStart(); err != nil {
		t.Fatal(err)
	}
	p, err := s.openHPass()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(p.close)
	return s, p
}
