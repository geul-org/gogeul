//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what wordsUint — up to 4 words, high first, as an unsigned integer
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// wordsUint reads ws (at most 4 words, high first) as an unsigned integer.
func wordsUint(ws []core.Word) uint64 {
	var n uint64
	for _, w := range ws {
		n = n<<16 | uint64(w)
	}
	return n
}
