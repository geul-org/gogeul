//ff:func feature=resolve type=helper control=sequence
//ff:what TestWordsUint — words high first as an integer; none is 0
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestWordsUint(t *testing.T) {
	if wordsUint([]core.Word{0xffff, 0xffff, 0xffff, 0xfffe}) != 0xfffffffffffffffe || wordsUint([]core.Word{0, 0x0980}) != 0x0980 || wordsUint(nil) != 0 {
		t.Fatal()
	}
}
