//ff:func feature=delta type=helper control=sequence
//ff:what edgeKey — the content key of a statement: its codec and the words after its Edge Id
package delta

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// edgeKey identifies a statement by content: two Edges of the same codec
// whose payloads match after the Edge Id are the same statement.
func edgeKey(codec uint16, rest []core.Word) string {
	return fmt.Sprintf("%02x:%04x", codec, rest)
}
