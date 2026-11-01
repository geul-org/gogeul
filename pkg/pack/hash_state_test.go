//ff:func feature=pack type=helper control=sequence
//ff:what TestHashState — no hash gives no state; a saved state stands for the same digest
package pack

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestHashState(t *testing.T) {
	if hashState(nil) != nil {
		t.Fatal("nil hash: want nil")
	}
	h := sha256.New()
	h.Write([]byte("abc"))
	got, err := sumOfState(hashState(h))
	want := sha256.Sum256([]byte("abc"))
	if err != nil || !bytes.Equal(got, want[:]) {
		t.Fatalf("%x %v", got, err)
	}
}
