//ff:func feature=pack type=helper control=sequence
//ff:what TestSumOfState — a marshaled sha256 state gives its digest; bytes that are no state are ErrCorrupt
package pack

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"testing"
)

func TestSumOfState(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("part"))
	got, err := sumOfState(hashState(h))
	if err != nil || !bytes.Equal(got, h.Sum(nil)) {
		t.Fatalf("%x %v", got, err)
	}
	if _, err := sumOfState([]byte("no state")); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("bad state: %v", err)
	}
}
