//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderDataPacks — data pack numbers in OBJECTS order, index objects left out; none in an empty manifest
package pack

import (
	"slices"
	"testing"
)

func TestReaderDataPacks(t *testing.T) {
	m := &manifest{objects: []object{{kind: kindIndex}, {kind: kindData}, {kind: kindData}, {kind: kindIndex}, {kind: kindData}}}
	if got := newReader(memStore{}, Options{}, m, "").DataPacks(); !slices.Equal(got, []uint32{1, 2, 4}) {
		t.Fatal(got)
	}
	if got := newReader(memStore{}, Options{}, &manifest{}, "").DataPacks(); len(got) != 0 {
		t.Fatal(got)
	}
}
