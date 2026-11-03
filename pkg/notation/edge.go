//ff:type feature=notation type=model
//ff:what Edge — an Edge read from a line, before its Edge Id is assigned: its #label, codec and payload after the Edge Id
package notation

import "github.com/geul-org/gogeul/pkg/core"

// Edge is a line read by ParseEdge. Label is the text of the line's leading
// `#label` ("" when none; a raw Edge's own Edge Id in decimal): a name for the
// Edge within a batch, not its Edge Id, which the writer assigns
// (delta.LocalIDs). Rest is the payload after the Edge Id.
type Edge struct {
	Label string
	Codec uint16
	Rest  []core.Word
}
