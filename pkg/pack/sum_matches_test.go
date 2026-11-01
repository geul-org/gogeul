//ff:func feature=pack type=helper control=sequence
//ff:what TestSumMatches — the bytes after the header up to the size match their saved state; other bytes do not; a bad state and a failing read are errors
package pack

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestSumMatches(t *testing.T) {
	b := bytes.Repeat([]byte("0123456789"), 10)
	h := sha256.New()
	h.Write(b[headerSize:90])
	open := streamOpen{Off: 90, Sum: hashState(h)}
	if ok, err := sumMatches(bytes.NewReader(b), open); !ok || err != nil {
		t.Fatalf("same bytes: %v", err)
	}
	b[70] = 'x'
	if ok, err := sumMatches(bytes.NewReader(b), open); ok || err != nil {
		t.Fatal("other bytes: want no match")
	}
	if _, err := sumMatches(bytes.NewReader(b), streamOpen{Off: 90, Sum: []byte("bad")}); err == nil {
		t.Fatal("bad state: want error")
	}
	if _, err := sumMatches(&badFile{readAt: true}, open); err == nil {
		t.Fatal("failing read: want error")
	}
}
