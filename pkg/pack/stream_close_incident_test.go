//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamCloseIncident — the 2026-11-01 stop: a 6c02e78 checkpoint with the data pack open, then Close sealed that pack and an index object and left a later index.part and a stray run before it was killed; the resume restores both open objects from o/ and writes the files of an unstopped run
package pack

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamCloseIncident(t *testing.T) {
	ref := streamCrashRef(t)
	found := false
	for n := 0; n < 200 && !found; n++ {
		out := filepath.Join(t.TempDir(), "set")
		s, old := streamCrashStart(t, out, true)
		streamCrashRest(t, s)
		left := n
		s.o.Build.Codec = crashCodec{inner: s.o.Build.Codec, left: &left}
		if _, err := s.Close(); err == nil {
			break
		}
		os.WriteFile(filepath.Join(out, streamDir, streamStateFile), old, 0o644) // the old Close saved nothing
		if found = streamIncidentState(t, out, old); !found {
			continue
		}
		if got := streamCrashResume(t, out); !maps.Equal(got, ref) {
			t.Fatalf("stop at compression %d: files differ", n)
		}
	}
	if !found {
		t.Fatal("no stop point reproduces the incident")
	}
}
