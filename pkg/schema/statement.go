//ff:type feature=schema type=model
//ff:what statement — a TRIPLE or QUALIFIER cut by its fixed layout: Edge Id, three records, the words left
package schema

import "github.com/geul-org/gogeul/pkg/core"

// statement is a TRIPLE (Predicate, Subject, Object; Rest = Rank) or a
// QUALIFIER (Property, Target, Value; Rest empty) of a BODY, cut by the
// layout the grammar fixes for these direct codes (edge/TRIPLE, QUALIFIER).
type statement struct {
	R     rec
	Codec uint16
	ID    uint64
	Parts [3][]core.Word
	Rest  []core.Word
}
