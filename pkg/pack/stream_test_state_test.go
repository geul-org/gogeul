//ff:func feature=pack type=writer control=sequence
//ff:what streamTestState — a saved state decoded from its gob bytes
package pack

import (
	"bytes"
	"encoding/gob"
	"testing"
)

// streamTestState decodes state bytes as ResumeStream does.
func streamTestState(t *testing.T, b []byte) streamState {
	t.Helper()
	var st streamState
	if err := gob.NewDecoder(bytes.NewReader(b)).Decode(&st); err != nil {
		t.Fatal(err)
	}
	return st
}
