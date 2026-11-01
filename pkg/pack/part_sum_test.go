//ff:func feature=pack type=helper control=sequence
//ff:what TestPartSum — the hash covers the bytes after the header up to the size; a failing read is an error
package pack

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestPartSum(t *testing.T) {
	b := bytes.Repeat([]byte("0123456789"), 10)
	h, err := partSum(bytes.NewReader(b), 80)
	want := sha256.Sum256(b[headerSize:80])
	if err != nil || !bytes.Equal(h.Sum(nil), want[:]) {
		t.Fatalf("%v", err)
	}
	if _, err := partSum(&badFile{readAt: true}, 80); err == nil {
		t.Fatal("failing read: want error")
	}
}
