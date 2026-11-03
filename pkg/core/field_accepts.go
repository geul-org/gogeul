//ff:func feature=core type=helper control=selection
//ff:what FieldAccepts — whether a record fits a FIELD's CODEC targets (slot markers only by exact WORD)
package core

import "slices"

// FieldAccepts reports whether r may fill a record field whose CODEC targets
// are ts (GEUL edge/CODEC). A Reference of a slot marker code (C < 3) is
// accepted only as a one-word slot marker (IsSlotMarker) whose exact WORD the
// list names: a field without CODECs ("any record") and wildcards never take
// one. Any other record fits an empty list, or a list with a target it matches
// (CodecMatches).
func FieldAccepts(ts []Word, r Record) bool {
	marker := r.Kind == KindReference && IsSlotMarkerCodec(r.Codec)
	switch {
	case marker && !IsSlotMarker(r):
		return false
	case len(ts) == 0:
		return !marker
	default:
		return slices.ContainsFunc(ts, func(t Word) bool { return CodecMatches(t, r.Kind, r.Codec) })
	}
}
