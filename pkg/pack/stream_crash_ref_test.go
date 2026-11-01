//ff:func feature=pack type=writer control=sequence
//ff:what streamCrashRef — the files of the crash test stream written without a stop
package pack

import (
	"path/filepath"
	"testing"
)

// streamCrashRef writes the crash test stream (documents 1–40, checkpoint
// at 25, the named record) and closes it without a stop.
func streamCrashRef(t *testing.T) map[string]string {
	t.Helper()
	out := filepath.Join(t.TempDir(), "ref")
	s, _ := streamCrashStart(t, out, false)
	streamCrashRest(t, s)
	if _, err := s.Close(); err != nil {
		t.Fatal(err)
	}
	return streamTestTree(t, out)
}
