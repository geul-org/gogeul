//ff:func feature=pack type=writer control=sequence
//ff:what streamCrashResume — resume a stopped crash test stream at its checkpoint "25", put the rest again and close it, returning the set's files
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

// streamCrashResume resumes out with streamCrashOptions (a fresh Zstd), as
// the caller would: from meta "25" it puts documents 26–40 and the named
// record again, then closes. The working folder must be gone after.
func streamCrashResume(t *testing.T, out string) map[string]string {
	t.Helper()
	s, err := ResumeStream(out, streamCrashOptions())
	if err != nil {
		t.Fatalf("resume: %v", err)
	}
	if string(s.Meta()) != "25" {
		t.Fatalf("meta %q", s.Meta())
	}
	streamCrashRest(t, s)
	if _, err := s.Close(); err != nil {
		t.Fatalf("close after resume: %v", err)
	}
	if _, err := os.Stat(filepath.Join(out, streamDir)); err == nil {
		t.Fatal("working folder left")
	}
	return streamTestTree(t, out)
}
