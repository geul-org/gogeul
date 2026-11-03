//ff:func feature=resolve type=helper control=sequence
//ff:what TestUintWords4 — a 64-bit document ID as 4 words, high first, and back through wordsUint
package resolve

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestUintWords4(t *testing.T) {
	if got := uintWords4(0x0001000200030004); !slices.Equal(got, []core.Word{1, 2, 3, 4}) || wordsUint(got) != 0x0001000200030004 {
		t.Fatal(got)
	}
}
