//ff:func feature=pack type=writer control=iteration dimension=2
//ff:what TestResumeStreamSealed — a stream killed while putting after its checkpoint, after packs were cut (part files gone or holding later objects), resumes from a hash-saving or a 6c02e78 checkpoint and writes the files of an unstopped run
package pack

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

func TestResumeStreamSealed(t *testing.T) {
	ref := streamCrashRef(t)
	for _, legacy := range []bool{false, true} {
		cut := 0
		for k := uint64(26); k <= 40; k++ {
			out := filepath.Join(t.TempDir(), "set")
			s, _ := streamCrashStart(t, out, legacy)
			streamTestPut(t, s, 26, k)
			b, _ := os.ReadFile(filepath.Join(out, streamDir, streamStateFile))
			r := &Stream{out: out, st: streamTestState(t, b)}
			if names, _ := r.unknownObjects(kindData); len(names) > 0 {
				cut++
			}
			if got := streamCrashResume(t, out); !maps.Equal(got, ref) {
				t.Fatalf("legacy %v, killed after document %d: files differ", legacy, k)
			}
		}
		if cut == 0 {
			t.Fatalf("legacy %v: no data pack was cut after the checkpoint", legacy)
		}
	}
}
