//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what hexToken — words as one token: 0x and four upper-case hex digits per word
package notation

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// hexToken writes ws as `0x` followed by 4 hex digits per word.
func hexToken(ws []core.Word) string {
	var b strings.Builder
	b.WriteString("0x")
	for _, w := range ws {
		fmt.Fprintf(&b, "%04X", uint16(w))
	}
	return b.String()
}
