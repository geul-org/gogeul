//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestMergeRunsPace — a long merge calls pace every paceEntries entries and stops at its error
package pack

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestMergeRunsPace(t *testing.T) {
	s := streamTestNew(t, nil)
	var b []byte
	for i := 0; i < paceEntries+1; i++ {
		var e hEntry
		e.key[0], e.key[1], e.key[2] = byte(i>>16), byte(i>>8), byte(i)
		b = appendHEntry(b, e)
	}
	if err := os.WriteFile(filepath.Join(s.out, streamDir, "big"), b, 0o644); err != nil {
		t.Fatal(err)
	}
	calls := 0
	pace := func() error { calls++; return nil }
	if err := mergeRuns(s.out, []string{"big"}, "m", pace); err != nil || calls != 1 {
		t.Fatalf("%d calls: %v", calls, err)
	}
	stop := errors.New("stop")
	if err := mergeRuns(s.out, []string{"big"}, "n", func() error { return stop }); !errors.Is(err, stop) {
		t.Fatal(err)
	}
}
