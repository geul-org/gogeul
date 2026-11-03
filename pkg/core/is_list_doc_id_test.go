//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsListDocID — the top three 4-word document IDs (issuers, docs, codecs) are reserved; 0 and others are not
package core

import "testing"

func TestIsListDocID(t *testing.T) {
	for d, want := range map[uint64]bool{0xFFFFFFFFFFFFFFFF: true, 0xFFFFFFFFFFFFFFFE: true, 0xFFFFFFFFFFFFFFFD: true, 0xFFFFFFFFFFFFFFFC: false, 0: false, 1: false} {
		if IsListDocID(d) != want {
			t.Fatalf("%016X", d)
		}
	}
}
