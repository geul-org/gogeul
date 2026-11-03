//ff:func feature=delta type=encoder control=sequence
//ff:what LocalIDs.Edge — the Edge Id for a statement: the previous ID of an equal statement, else the next new ID
package delta

import "github.com/geul-org/gogeul/pkg/core"

// Edge returns the Edge Id for an Edge of codec whose payload after the Edge
// Id is rest (records and raw words as written): the ID an equal statement
// had in the previous version (equal statements in turn, so duplicates keep
// their order), else a new ID from the auto increment. The caller then
// writes the Edge with that ID.
func (l *LocalIDs) Edge(codec uint16, rest []core.Word) uint64 {
	k := edgeKey(codec, rest)
	if ids := l.edges[k]; len(ids) > 0 {
		l.edges[k] = ids[1:]
		l.used[ids[0]] = true
		return ids[0]
	}
	return l.fresh()
}
