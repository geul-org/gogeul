//ff:func feature=pack type=writer control=sequence
//ff:what streamTestLegacy — rewrite a stream's saved state without hash states, as gogeul 6c02e78 saved it, and return its bytes
package pack

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
	"testing"
)

// streamTestLegacy clears Data.Sum and Index.Sum in out's state file (gob
// leaves zero fields out, so the bytes are what a checkpoint without the
// field wrote).
func streamTestLegacy(t *testing.T, out string) []byte {
	t.Helper()
	path := filepath.Join(out, streamDir, streamStateFile)
	b, err := os.ReadFile(path)
	var st streamState
	if err == nil {
		err = gob.NewDecoder(bytes.NewReader(b)).Decode(&st)
	}
	if err != nil {
		t.Fatal(err)
	}
	st.Data.Sum, st.Index.Sum = nil, nil
	var w bytes.Buffer
	gob.NewEncoder(&w).Encode(st)
	if err := os.WriteFile(path, w.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	return w.Bytes()
}
