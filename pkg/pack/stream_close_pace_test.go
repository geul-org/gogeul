//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamClosePace — a Close stopped by Pace at each of its safe points (every phase met) returns ErrStopped and Pace's error; ResumeStream and Close then write the files of a Close never stopped
package pack

import (
	"errors"
	"maps"
	"path/filepath"
	"testing"
)

func TestStreamClosePace(t *testing.T) {
	ref := streamCrashRef(t)
	stop := errors.New("low memory")
	seen := map[string]bool{}
	n := 0
	for ; ; n++ {
		out := filepath.Join(t.TempDir(), "set")
		err := streamPacedClose(t, out, n, stop, seen)
		if err == nil {
			break
		}
		if !errors.Is(err, ErrStopped) || !errors.Is(err, stop) {
			t.Fatalf("point %d: %v", n, err)
		}
		if got := streamCrashResume(t, out); !maps.Equal(got, ref) {
			t.Fatalf("stop at point %d: files differ", n)
		}
	}
	if n < 8 {
		t.Fatalf("only %d safe points", n)
	}
	for _, phase := range []string{"close:data", "close:p", "close:merge", "close:h", "close:tree", "close:index"} {
		if !seen[phase] {
			t.Fatalf("phase %s never met: %v", phase, seen)
		}
	}
}
