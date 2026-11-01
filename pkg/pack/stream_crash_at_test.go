//ff:func feature=pack type=writer control=sequence
//ff:what streamCrashAt — write the crash test stream with Close stopped after n compressions, resume it (stopping it once more when asked) and check its files equal the unstopped run's; false when Close did not stop
package pack

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

// streamCrashAt runs one stop point: the stream starts (legacy: from a
// hash-less checkpoint, whose state is put back after the stop, as the old
// Close saved nothing), Close stops at compression n, and the resumed run
// must write ref. With twice the resumed Close is stopped after one
// compression first and resumed again. It returns false when Close
// finished without stopping (n is past the last compression); the files
// must then equal ref too.
func streamCrashAt(t *testing.T, ref map[string]string, legacy, twice bool, n int) bool {
	t.Helper()
	out := filepath.Join(t.TempDir(), "set")
	s, old := streamCrashStart(t, out, legacy)
	streamCrashRest(t, s)
	left := n
	s.o.Build.Codec = crashCodec{inner: s.o.Build.Codec, left: &left}
	if _, err := s.Close(); err == nil {
		if !maps.Equal(streamTestTree(t, out), ref) {
			t.Fatalf("legacy %v: unstopped run differs", legacy)
		}
		return false
	}
	if legacy {
		os.WriteFile(filepath.Join(out, streamDir, streamStateFile), old, 0o644)
	}
	if twice {
		again, err := ResumeStream(out, streamCrashOptions())
		one := 1
		if err == nil {
			again.o.Build.Codec = crashCodec{inner: again.o.Build.Codec, left: &one}
			_, err = again.Close()
		}
		if err == nil {
			t.Fatalf("n %d: second stop did not stop", n)
		}
	}
	if got := streamCrashResume(t, out); !maps.Equal(got, ref) {
		t.Fatalf("legacy %v, stop at compression %d: files differ", legacy, n)
	}
	return true
}
