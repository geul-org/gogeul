//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestMergeRuns — runs merge into one sorted run keeping every entry; a missing run, an uncreatable output, a cut run and a full disk are errors
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMergeRuns(t *testing.T) {
	s := streamTestNew(t, nil)
	names := streamTestRuns(t, s, []byte{1, 5, 9}, []byte{2, 5}, []byte{7})
	if err := mergeRuns(s.out, names, "merged", nil); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(filepath.Join(s.out, streamDir, "merged"))
	var keys []byte
	for i := 0; i < len(b); i += hEntrySize {
		keys = append(keys, parseHEntry(b[i:i+hEntrySize]).key[0])
	}
	if string(keys) != string([]byte{1, 2, 5, 5, 7, 9}) || parseHEntry(b[2*hEntrySize:]).loc.Off != 50 {
		t.Fatalf("merged keys %v", keys)
	}
	if err := mergeRuns(s.out, []string{"run-missing"}, "x", nil); err == nil {
		t.Fatal("missing run: want error")
	}
	if err := mergeRuns(s.out, names, "no/x", nil); err == nil {
		t.Fatal("uncreatable output: want error")
	}
	cut := filepath.Join(s.out, streamDir, names[0])
	good, _ := os.ReadFile(cut)
	os.WriteFile(cut, good[:hEntrySize+3], 0o644)
	if err := mergeRuns(s.out, names, "y", nil); err == nil {
		t.Fatal("cut run: want error")
	}
	os.WriteFile(cut, good, 0o644)
	os.Symlink("/dev/full", filepath.Join(s.out, streamDir, "full"))
	if err := mergeRuns(s.out, names, "full", nil); err == nil {
		t.Fatal("full disk: want error")
	}
}
