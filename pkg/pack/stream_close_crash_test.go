//ff:func feature=pack type=writer control=iteration dimension=2
//ff:what TestStreamCloseCrash — a Close stopped at every compression (with a hash-saving checkpoint, sometimes stopped again after the resume; and with a 6c02e78 hash-less checkpoint, left as the old Close left it), at the manifest and at latest.txt, resumed, writes the same files as a Close never stopped
package pack

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamCloseCrash(t *testing.T) {
	ref := streamCrashRef(t)
	for _, legacy := range []bool{false, true} {
		points := 0
		for n := 0; streamCrashAt(t, ref, legacy, !legacy && n%3 == 1, n); n++ {
			points++
		}
		t.Logf("legacy %v: %d stop points", legacy, points)
		if points < 20 {
			t.Fatalf("legacy %v: only %d stop points", legacy, points)
		}
	}
	for name, breakIt := range map[string]func(out string) func(){
		"manifest":   func(out string) func() { m := filepath.Join(out, "m"); os.RemoveAll(m); os.WriteFile(m, nil, 0o644); return func() { os.Remove(m); os.Mkdir(m, 0o755) } },
		"latest.txt": func(out string) func() { l := filepath.Join(out, latestFile, "x"); os.MkdirAll(l, 0o755); return func() { os.RemoveAll(filepath.Dir(l)) } },
	} {
		out := filepath.Join(t.TempDir(), "set")
		s, _ := streamCrashStart(t, out, false)
		streamCrashRest(t, s)
		fix := breakIt(out)
		if _, err := s.Close(); err == nil {
			t.Fatalf("%s: want error", name)
		}
		fix()
		if got := streamCrashResume(t, out); !maps.Equal(got, ref) {
			t.Fatalf("stop at %s: files differ", name)
		}
	}
}
