//ff:func feature=core type=helper control=selection
//ff:what CodeName — the name of one TAG code as its codec document binds it (§1.1 124): 5Q <체계>-<양화>, even/odd pair <이름>-EVEN·-ODD, slot marker name; otherwise CodecName
package core

// quantifierCodeNames spell the five Quantifiers (QuantifierName order) in
// code names.
var quantifierCodeNames = [...]string{"IDENTIFIED", "GENERIC", "FREE-CHOICE", "UNIVERSAL", "VARIABLE"}

// pairCodeSuffix is the suffix of the even (0) and odd (1) code of a Value
// pair.
var pairCodeSuffix = [...]string{"-EVEN", "-ODD"}

// CodeName returns the name the codec document of code k/codec gives its
// DECLARE (migration.md §1.1 124): each 5Q code is its scheme and Quantifier
// (LOCAL-GENERIC, EXTERNAL-POINTER-1WORD-VARIABLE), each code of an even/odd
// Value pair (two codes CodecName names alike) the pair name with -EVEN or
// -ODD (UTF8-TEXT-ODD), a slot marker its marker name (ABSENT). Every other
// code is named as by CodecName, which keeps naming the scheme and the pair
// for records and notation. Unassigned codes return "".
func CodeName(k Kind, codec uint16) string {
	n := CodecName(k, codec)
	switch {
	case k == KindReference && IsSlotMarkerCodec(codec):
		return SlotMarkerName(codec)
	case k == KindReference && n != "":
		if q, ok := RefQuantifier(codec); ok {
			return n + "-" + quantifierCodeNames[q]
		}
		return n
	case k == KindValue && n != "" && valueNames[codec^1] == n:
		return n + pairCodeSuffix[codec&1]
	default:
		return n
	}
}
