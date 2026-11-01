//ff:func feature=pack type=writer control=sequence
//ff:what TestHPassClose — the runs and the head file are closed
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHPassClose(t *testing.T) {
	dir := t.TempDir()
	run, heads := filepath.Join(dir, "run"), filepath.Join(dir, "heads")
	os.WriteFile(run, nil, 0o644)
	os.WriteFile(heads, nil, 0o644)
	rf, _ := os.Open(run)
	hf, _ := os.Open(heads)
	(&hPass{rs: []*runReader{{f: rf}}, heads: hf}).close()
	if _, err := rf.Stat(); err == nil {
		t.Fatal("run left open")
	}
	if _, err := hf.Stat(); err == nil {
		t.Fatal("head file left open")
	}
}
