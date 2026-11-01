//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamHStart — an empty head file and a zero position per run are saved with phase closeHPass; an unwritable head file is an error
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamHStart(t *testing.T) {
	s := streamTestNew(t, nil)
	s.st.Runs = []string{"a", "b"}
	s.st.Close.HHeads = 3
	heads := filepath.Join(s.out, streamDir, streamHeadsFile)
	os.WriteFile(heads, []byte("old"), 0o644)
	if err := s.hStart(); err != nil || s.st.Close.Phase != closeHPass || len(s.st.Close.Taken) != 2 || s.st.Close.HHeads != 0 {
		t.Fatal(err)
	}
	if fi, _ := os.Stat(heads); fi.Size() != 0 {
		t.Fatal("head file not emptied")
	}
	os.Remove(heads)
	os.MkdirAll(filepath.Join(heads, "x"), 0o755)
	if err := s.hStart(); err == nil {
		t.Fatal("unwritable head file: want error")
	}
}
