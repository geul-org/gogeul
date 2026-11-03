//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestPublicationTail — tail lengths per codec; other codecs are not version list records
package core

import "testing"

func TestPublicationTail(t *testing.T) {
	for c, want := range map[uint16]int{EdgePublication: 16, EdgeDeprecation: 0, EdgeTombstone: 1} {
		if n, ok := publicationTail(c); !ok || n != want {
			t.Fatalf("0x%02X: %d %v", c, n, ok)
		}
	}
	if _, ok := publicationTail(0x00); ok {
		t.Fatal("0x00: want false")
	}
}
