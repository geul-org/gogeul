//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestSamplesBySum — samples come back in SHA-256 order as a permutation, so alternating Put order (list, version, …) no longer sends one kind to the training side
package pack

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestSamplesBySum(t *testing.T) {
	var in [][]byte
	for i := 0; i < 64; i++ {
		in = append(in, []byte(fmt.Sprintf("versions.geul list %d", i)), []byte(fmt.Sprintf("version document %d with a longer body", i)))
	}
	out := samplesBySum(in)
	if len(out) != len(in) {
		t.Fatalf("len %d", len(out))
	}
	lists := 0
	for i, b := range out {
		a, c := sha256.Sum256(out[max(i-1, 0)]), sha256.Sum256(b)
		if bytes.Compare(a[:], c[:]) > 0 {
			t.Fatalf("not in SHA-256 order at %d", i)
		}
		if i%2 == 0 && bytes.HasPrefix(b, []byte("versions.geul")) {
			lists++
		}
	}
	if lists == 0 || lists == 64 {
		t.Fatalf("even (training) side holds %d of 64 lists: one kind only", lists)
	}
	if len(samplesBySum(nil)) != 0 {
		t.Fatal("nil")
	}
}
